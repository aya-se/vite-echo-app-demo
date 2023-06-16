<script setup>
import { store } from "../store.js";
import { ref } from "vue";

const input = ref("");
</script>
<template>
  <div>Edit</div>
  <TransitionGroup name="list">
    <div v-for="item in store.globalList" :key="item">
      <p>{{ item }}</p>
    </div>
  </TransitionGroup>
  <button @click="store.count++">Increment</button>
  <input v-model="input" />
  <button @click="store.globalList.push(input)">Add</button>
</template>
<style>
.list-enter-active,
.list-leave-active {
  transition: all 0.5s;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>
