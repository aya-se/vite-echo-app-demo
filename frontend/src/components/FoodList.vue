<script setup lang="ts">
import { ref } from "vue";
const API_URL = "http://localhost:1323";
interface Item {
  name: string;
  price: number;
}
const items = ref<Item[]>([]);

fetch(`${API_URL}/foods`)
  .then((res) => res.json())
  .then((data) => {
    items.value = data;
  });
</script>

<template>
  <div>
    <div>FoodList</div>
    <ul>
      <li v-for="item in items" :key="item.name">
        <div>名前: {{ item.name }}</div>
        <div>価格: {{ item.price }}円</div>
      </li>
    </ul>
  </div>
</template>
<style lang="scss"></style>
