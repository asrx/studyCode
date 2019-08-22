<template>
  <form v-on:submit.prevent="addTodo(newTodo)">
  	<div class="form-group">
  		<input type="text"
  			v-model="newTodo.title" class="form-control" placeholder="Add a new Todo"/>
  	</div>
  	<div class="form-group">
  		<input type="submit" class="btn btn-success" value="Add Todo"/>
  	</div>
  </form>
</template>

<script>
  export default{
    computed: {
      newTodo(){
        return this.$store.state.newTodo
      }
    },
    methods: {
    	addTodo(newTodo){
    		this.$store.dispatch('insertTodo',newTodo)
    	}
    }
  }
</script>

<style>
</style>
