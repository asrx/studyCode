// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Vuex from 'vuex'

import App from './App'
import router from './router'

import store from './vuex/store.js'

Vue.use(VueAxios,axios)
Vue.use(Vuex)

Vue.config.productionTip = false

Vue.prototype.baseUrl = 'http://longyuan-query.test'




/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  router,
  components: { App },
  template: '<App/>'
})
