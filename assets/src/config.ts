import global from './config/global.json'; // 使用 TypeScript 的 JSON 导入功能

// 定义 URI 类型，使用索引签名
interface URI {
    [key: string]: string; // 允许任意字符串键，值为字符串
}

interface HOST {
    [key: string]: string; // 允许任意字符串键，值为字符串
}

// 定义 Config 类型
interface Config {
    HOST: HOST;
    URI: URI;
    REFRESH_INTERVAL: number;
}

// JSON 数据
export const config: Config = global
console.log("global config: " + config)

// 获取指定的URI
export function wisdom_url(uri: keyof URI): string {
    const host = config.HOST.Wisdom; // 从配置中获取 HOST
    const endpoint = config.URI[uri]; // 从配置中获取 URI
    if (!endpoint) {
        throw new Error(`URI ${uri} not found in config`);
    }
    return `${host}${endpoint}`; // 返回完整的 URL
}

const randomWisdomUrl = wisdom_url("GetRandomWisdom");
console.log("Random Wisdom URL:", randomWisdomUrl); // 输出: http://127.0.0.1:1666/api/wisdom?random=1

const saveWisdomUrl = wisdom_url("SaveWisdom");
console.log("Save Wisdom URL:", saveWisdomUrl); // 输出: http://127.0.0.1:1666/api/wisdom
