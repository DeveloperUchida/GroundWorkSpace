document.addEventListener("DOMContentLoaded", () => {
    // 外部HTMLファイルを読み込む
    fetch("time.html")
        .then(response => {
            if (!response.ok) {
                throw new Error("HTMLファイルの読み込みに失敗しました。");
            }
            return response.text();
        })
        .then(htmlContent => {
            // 読み込んだHTMLを挿入
            document.getElementById("content").innerHTML = htmlContent;
        })
        .catch(error => {
            console.error(error);
            document.getElementById("content").innerHTML = `
                <div class="maintenance-message">
                    <h1>エラー</h1>
                    <p>調整中のメッセージを読み込めませんでした。</p>
                </div>
            `;
        });
});