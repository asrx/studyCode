<template>
  <div class="todo">
    <div class="loading" v-if="loading">loading...</div>
    <div class="error" v-if="error">{{error}}</div>
    <div class="content" v-if="todo">
      <h2>{{todo.title}}</h2>
    </div>
  </div>

</template>

<script>
  export default{
    data () {
        return {
          loading: false,
          todo: null,
          error: null
        }
      },
      created () {
        // fetch the data when the view is created and the data is
        // already being observed
        this.fetchData()
      },
      watch: {
        // call again the method if the route changes
        '$route': 'fetchData'
      },
      methods: {
        fetchData () {
          this.error = this.todo = null
          this.loading = true
          // replace `getPost` with your data fetching util / API wrapper
          this.axios.get('http://longyuan-query.test/api/todo/' + this.$route.params.id).then(response=>{
            this.loading = false
            console.log(response)
            this.todo = response.data
          })
        }
      }
  }
</script>

<style>

</style>
