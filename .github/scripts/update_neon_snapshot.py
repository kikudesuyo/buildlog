"""Replace the manual Neon snapshot with a fresh snapshot."""

from __future__ import annotations

import json
import os
import sys
import time
from dataclasses import dataclass
from urllib.error import HTTPError, URLError
from urllib.parse import urlencode
from urllib.request import Request, urlopen


@dataclass(frozen=True)
class Config:
    api_key: str
    api_url: str
    project_id: str


def request_json(
    config: Config,
    url: str,
    method: str = "GET",
    retries: int = 0,
) -> dict:
    request = Request(
        url,
        headers={
            "Accept": "application/json",
            "Authorization": f"Bearer {config.api_key}",
        },
        method=method,
    )

    for attempt in range(retries + 1):
        try:
            with urlopen(request, timeout=30) as response:
                body = response.read()
            return json.loads(body) if body else {}
        except (HTTPError, URLError, TimeoutError) as error:
            if attempt == retries:
                raise RuntimeError(f"Neon API request failed: {method} {url}") from error
            time.sleep(2**attempt)

    raise AssertionError("unreachable")


def load_config() -> Config:
    values = {
        "NEON_API_KEY": os.environ.get("NEON_API_KEY", ""),
        "NEON_API_URL": os.environ.get("NEON_API_URL", ""),
        "NEON_PROJECT_ID": os.environ.get("NEON_PROJECT_ID", ""),
    }
    missing = [name for name, value in values.items() if not value]
    if missing:
        raise RuntimeError(f"Missing required environment variable(s): {', '.join(missing)}")

    return Config(
        api_key=values["NEON_API_KEY"],
        api_url=values["NEON_API_URL"].rstrip("/"),
        project_id=values["NEON_PROJECT_ID"],
    )


def replace_snapshot(config: Config) -> None:
    snapshots_url = f"{config.api_url}/projects/{config.project_id}/snapshots"
    snapshots = request_json(config, snapshots_url, retries=3)

    for snapshot in snapshots.get("snapshots", []):
        snapshot_id = snapshot.get("id")
        if not snapshot_id:
            continue
        print(f"Deleting snapshot: {snapshot_id}")
        request_json(config, f"{snapshots_url}/{snapshot_id}", method="DELETE")

    for attempt in range(12):
        snapshots = request_json(config, snapshots_url, retries=3)
        if not snapshots.get("snapshots"):
            break
        if attempt == 11:
            raise RuntimeError("The previous Neon snapshot was not deleted in time.")
        time.sleep(10)

    branches_url = f"{config.api_url}/projects/{config.project_id}/branches?limit=100"
    branches = request_json(config, branches_url, retries=3)
    branch_id = next(
        (
            branch.get("id")
            for branch in branches.get("branches", [])
            if branch.get("parent_id") is None
        ),
        None,
    )
    if not branch_id:
        raise RuntimeError("Could not find the Neon root branch.")

    snapshot_name = f"buildlog-{time.strftime('%Y%m%d', time.gmtime())}"
    snapshot_url = (
        f"{config.api_url}/projects/{config.project_id}/branches/{branch_id}/snapshot?"
        f"{urlencode({'name': snapshot_name})}"
    )
    print(f"Creating snapshot: {snapshot_name}")
    request_json(config, snapshot_url, method="POST")
    print("Snapshot created successfully.")


if __name__ == "__main__":
    try:
        replace_snapshot(load_config())
    except RuntimeError as error:
        print(error, file=sys.stderr)
        sys.exit(1)
