import './css/index.css';

// @ts-ignore
import image from './imgs/code.png'; // 图片
import {initializeConfig} from './config' // 异步加载config代码
import {autoRefreshWisdom} from './autoload' // 自动刷新

// 主函数
function main() {
    console.log("wisdom amazing!")

    // 初始化配置，然后自动刷新
    initializeConfig().then(function () {
        autoRefreshWisdom()
    })

    // 图片
    const img = document.createElement('img');
    img.src = image;
    document.body.appendChild(img);
}

main()

// 热加载
if (module.hot) {
    module.hot.accept('/dist/bundle.js', function() {
        console.log('Updating module...');
        // Handle module update logic here
    });
}
