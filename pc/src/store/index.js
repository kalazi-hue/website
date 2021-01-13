// 1、引入vuex
import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

//2、定义公共的数据及方法
const store = new Vuex.Store({
    state:{
        adList: [],
        configList: []
    },
    mutations: {
        editAdList(state, newVal){
            state.adList = newVal
        },
        editConfigList(state, newVal){
            state.configList = newVal
        }
    }
})
///3、导出
export default store