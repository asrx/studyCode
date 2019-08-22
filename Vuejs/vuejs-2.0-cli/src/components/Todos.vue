<template>
  <div id="todos">
    <ul class="list-group">
      <li class="list-group-item" v-for="(todo,index) in todos" v-bind:class="{ 'completed' : todo.completed }">

        <router-link :to="{name:'todo',params:{id:todo.id}}">{{todo.title}}</router-link>

        <button class="btn btn-xs pull-right"
          v-bind:class="[todo.completed ? 'btn-danger' : 'btn-success']"
          v-on:click="toggleCompletion(todo)">
          {{ todo.completed ? 'undo' : 'completed' }}
        </button>

        <button class="btn btn-warning btn-xs pull-right"
          v-on:click="deleteTodo(index,todo)">
          Delete
        </button>
      </li>
    </ul>

    <todo-form></todo-form>
  </div>

</template>

<script>
  import TodoForm from '@/components/TodoForm'
  export default{
    name: 'todos',

    mounted() {
      this.$store.dispatch('getTodos')
    },
    methods:{
    	deleteTodo(index,todo){
    		// this.todos.splice(index,1)
        this.$store.dispatch('removeTodo',index, todo)
    	},
    	toggleCompletion(todo){
    		this.$store.dispatch('completeTodo', todo)
    	}
    },
    computed:{
      todos(){
        return this.$store.state.todos
      },

      todoCount(){
        return this.$store.state.todos.length
      }
    },
    components:{
      TodoForm
    }
  }
</script>

<style>
  .completed{
    color: #5cb85c;
    text-decoration: line-through;
  }
  .margin-right-10{
    margin-right: 10px;
  }
</style>
