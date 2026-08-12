import os
import unittest
from unittest.mock import patch

import update_neon_snapshot


class FakeResponse:
    def __init__(self, payload):
        self.payload = payload

    def __enter__(self):
        return self

    def __exit__(self, *_args):
        return False

    def read(self):
        import json

        return json.dumps(self.payload).encode()


class UpdateNeonSnapshotTest(unittest.TestCase):
    def test_replaces_existing_snapshot_from_root_branch(self):
        responses = iter(
            [
                {"snapshots": [{"id": "snapshot-1"}]},
                {},
                {"snapshots": []},
                {"branches": [{"id": "root-branch", "parent_id": None}]},
                {},
            ]
        )

        with (
            patch.dict(
                os.environ,
                {
                    "NEON_API_KEY": "test-key",
                    "NEON_API_URL": "https://console.neon.tech/api/v2",
                    "NEON_PROJECT_ID": "project-id",
                },
            ),
            patch.object(
                update_neon_snapshot,
                "urlopen",
                side_effect=lambda _request, timeout: FakeResponse(next(responses)),
            ) as mocked_urlopen,
            patch.object(update_neon_snapshot.time, "sleep"),
        ):
            update_neon_snapshot.replace_snapshot(update_neon_snapshot.load_config())

        requests = [call.args[0] for call in mocked_urlopen.call_args_list]
        self.assertEqual(requests[1].method, "DELETE")
        self.assertIn("name=buildlog-", requests[-1].full_url)
        self.assertEqual(len(requests), 5)


if __name__ == "__main__":
    unittest.main()
