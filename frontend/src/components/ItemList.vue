<script setup lang="ts">
import { ref } from "vue";

interface Item {
  name: string;
  price: number;
}
const newItemName = ref<string>("");
const newItemPrice = ref<number>(0);

const items = ref<Item[]>([
  { name: "たまご", price: 100 },
  { name: "りんご", price: 160 },
]);

const addItem = () => {
  if (newItemName.value === "" || newItemPrice.value === 0) {
    return;
  }
  items.value.push({ name: newItemName.value, price: newItemPrice.value });
  newItemName.value = "";
  newItemPrice.value = 0;
};
</script>

<template>
  <div>
    <div>ItemList</div>
    <ul>
      <li
        v-for="item in items"
        :key="item.name"
        :class="{ over500: item.price >= 500 }"
      >
        <div>名前: {{ item.name }}</div>
        <div>{{ item.price }} 円</div>
      </li>
    </ul>
  </div>
  <div>
    <label>名前<input v-model="newItemName" type="text" /></label>
    <label>価格<input v-model="newItemPrice" type="number" /></label>
    <button @click="addItem">追加</button>
  </div>
</template>

<style>
.over500 {
  color: red;
}
</style>
