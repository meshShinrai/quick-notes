<template>
  <div class="container">
    <h1>Quick Notes</h1>

    <form @submit.prevent="createNote">
      <input v-model="newNote" placeholder="Write a note..." required />
      <button type="submit">Add Note</button>
    </form>

    <ul>
      <li v-for="note in notes" :key="note.id">
        <p>{{ note.content }}</p>
        <button @click="deleteNote(note.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL

const notes = ref([])
const newNote = ref('')

const fetchNotes = async () => {
  const res = await axios.get(`${API_BASE_URL}/notes`)
  notes.value = res.data
}

const createNote = async () => {
  await axios.post(`${API_BASE_URL}/notes`, { content: newNote.value })
  newNote.value = ''
  await fetchNotes()
}

const deleteNote = async (id) => {
  await axios.delete(`${API_BASE_URL}/notes/${id}`)
  await fetchNotes()
}

onMounted(fetchNotes)
</script>

<style>
.container {
  max-width: 600px;
  margin: auto;
  font-family: sans-serif;
}
form {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}
input {
  flex: 1;
  padding: 8px;
}
button {
  padding: 8px 12px;
}
li {
  list-style: none;
  margin-bottom: 10px;
  border-bottom: 1px solid #ccc;
  padding-bottom: 5px;
}
</style>
