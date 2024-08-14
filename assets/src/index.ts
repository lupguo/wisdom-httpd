import {initializeConfig} from './config'
import {autoRefreshWisdom} from './autoload'

function main() {
    console.log('init config: '+initializeConfig())
    autoRefreshWisdom()
}

// 自动加载
console.log("wisdom amazing!" + initializeConfig())