import {config, wisdom_url} from './config'

// 发送请求以获取智慧内容
async function fetchWisdom() {
    let response = await fetch(wisdom_url("GetRandomWisdom"));
    if (!response.ok) {
        console.error(`HTTP error! Status: ${response.status}`);
        return null; // 返回 null 表示获取失败
    }
    return await response.json();
}


// 自动请求wisdomURI地址, 从数据中提取句子并更新模板
function refreshWisdom() {
    console.log("refreshWisdom...");

    fetchWisdom().then(
        function (data) {
            let cont = document.querySelector('.content')
            if (cont != null) {
                cont.textContent = data.sentence
            }
        }
    )
}

refreshWisdom()

setInterval(refreshWisdom, config.REFRESH_INTERVAL);
