import {initializeConfig} from './config'
import {autoRefreshWisdom} from './autoload'

function main() {
    console.log("wisdom amazing!")
    initializeConfig().then(function () {
        autoRefreshWisdom()
    })
}

main()