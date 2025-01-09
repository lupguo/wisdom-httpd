export interface Config {
    WISDOM_URI: string;
    REFRESH_INTERVAL: number;
}

// 定义全局变量
export let globalConfig: Config; // 初始为 null，表示尚未加载

// 在 JavaScript 中加载配置
async function loadConfig(): Promise<Config> {
    const response = await fetch('/dist/config/refresh.json'); // 确保路径正确
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    const data = await response.json();

    // 确保返回的数据符合 Config 类型
    return {
        WISDOM_URI: data.WISDOM_URI,
        REFRESH_INTERVAL: data.REFRESH_INTERVAL
    };
}

// 初始化刷新配置
export async function initializeConfig() {
    try {
        globalConfig = await loadConfig(); // 将加载的配置赋值给全局变量
        return "load succ"; // 返回成功信息
    } catch (error) {
        return `load config error: ${error}`; // 返回错误信息
    }
}

