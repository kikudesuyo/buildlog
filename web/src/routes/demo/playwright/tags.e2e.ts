import { test, expect } from "@playwright/test";

test("verify tagging and filtering logic and capture screenshots", async ({
  page,
}) => {
  const uniqueTitle = "E2Eタグテスト_" + Date.now();

  // 1. 管理画面に遷移して新規つぶやき作成
  await page.goto("http://localhost:5173/admin");
  await page.waitForLoadState("networkidle");

  await page.click("text=つぶやく");
  await page.waitForURL("**/admin/diary/new");

  // フォーム入力
  await page.fill('input[placeholder="タイトルを入力..."]', uniqueTitle);
  await page.fill(
    'textarea[placeholder="物語を書き始めましょう..."]',
    "これはタグ機能のE2Eテストです。",
  );

  // タグの追加
  // '+ タグを追加' ボタンをクリック
  page.on("dialog", async (dialog) => {
    expect(dialog.message()).toContain("タグ名を入力してください");
    await dialog.accept("GoTag");
  });
  await page.click("text=+ タグを追加");

  // 2つ目のタグ追加
  page.removeAllListeners("dialog");
  page.on("dialog", async (dialog) => {
    await dialog.accept("SvelteTag");
  });
  await page.click("text=+ タグを追加");

  // 保存 (投稿する)
  await page.click("text=投稿する");
  await page.waitForURL("**/admin");

  // 2. 一般向けつぶやきフィードへ
  await page.goto("http://localhost:5173/");
  await page.waitForLoadState("networkidle");

  const post = page.locator("article", { hasText: uniqueTitle }).first();
  await expect(post).toBeVisible();

  // タグバッジが表示されていることを確認
  const goTagBadge = post.locator('a:text-is("#GoTag")');
  const svelteTagBadge = post.locator('a:text-is("#SvelteTag")');
  await expect(goTagBadge).toBeVisible();
  await expect(svelteTagBadge).toBeVisible();

  // タグをクリックして絞り込む
  await goTagBadge.click();
  await page.waitForURL("**/?tag=GoTag");
  await page.waitForLoadState("networkidle");

  // タグフィルタのヘッダー表示を確認
  const filterHeader = page.locator("text=タグフィルタ:");
  await expect(filterHeader).toBeVisible();
  await expect(page.locator('span:text-is("GoTag")')).toBeVisible();

  // スクリーンショット撮影
  await page.screenshot({
    path: "static/screenshots/tags-filter.png",
    fullPage: true,
  });

  // フィルタを解除する
  await page.click('a[title="フィルタ解除"]');
  await page.waitForURL("http://localhost:5173/");
  await expect(filterHeader).not.toBeVisible();
});
