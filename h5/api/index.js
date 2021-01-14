// import axios from 'axios'
// import _ from 'lodash'

let configApi = ''

if (process.env.NODE_ENV === 'development') { // 测试
  // api = 'https://apitest.saohu00.com/v1'
} else { // 线上
  // api = 'https://www.8dafu.com/api'
}

// 测试地址
// api = 'https://apitest.saohu00.com/v1'

//线上地址
configApi = 'https://niuniutv5.com/api'
 
export default { configApi }
