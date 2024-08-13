const WISDOM_URI = 'http://127.0.0.1:1666/api/wisdom';
const REFRESH_INTERVAL = 10000; // 刷新间隔时间（毫秒）

// 发送请求以获取智慧内容
async function fetchWisdom() {
    try {
        const response = await fetch(WISDOM_URI);
        if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
        }
        const data = await response.json();
        console.log("Fetched wisdom:", data);
        return data;
    } catch (error) {
        console.error("Failed to fetch wisdom:", error);
    }
}

// 更新模板中的内容
function updateTemplate(sentence: string | null) {
    const contentDiv = document.querySelector('.content');
    if (contentDiv) {
        contentDiv.textContent = sentence; // 填充句子到模板
    } else {
        console.error("Content div not found!");
    }
}

// 自动请求wisdomURI地址
function refreshWisdom() {
    console.log("Refreshing wisdom...");
    fetchWisdom().then(function (data) {
        // 从数据中提取句子并更新模板
        if (data && data.status === "success" && data.data && data.data.PageData) {
            const sentence = data.data.PageData.sentence;
            updateTemplate(sentence);
        } else {
            console.error("Invalid data structure:", data);
        }
    }).catch(function (err) {
        console.error("Got error: " + err);
    });
}

// 启动自动刷新
export default function autoRefreshWisdom() {
    setInterval(refreshWisdom, REFRESH_INTERVAL);
}
