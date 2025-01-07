import {globalConfig} from './config'

// 发送请求以获取智慧内容
async function fetchWisdom() {
    try {
        const response = await fetch(globalConfig.WISDOM_URI);
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
        console.info("rsp data:", data);
        updateTemplate(data.sentence);

    }).catch(function (err) {
        console.error("Got error: " + err);
    });
}

// 基于配置参数，自动请求后台API接口
export function autoRefreshWisdom() {
    console.log("begin auto refresh wisdom...")
    setInterval(refreshWisdom, globalConfig.REFRESH_INTERVAL);
}

