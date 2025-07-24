<template>
  <div class="container">
    <h1>Quick Notes</h1>
    <Notes />
  </div>
</template>

<script setup>
import Notes from './components/NotesView.vue'
</script>

<style>
.container {
  max-width: 600px;
  margin: auto;
  padding: 1rem;
  font-family: sans-serif;
}
</style>
