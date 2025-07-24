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
  try {
    const res = await axios.get(`${API_BASE_URL}/notes`)
    notes.value = res.data
  } catch (error) {
    console.error('Error fetching notes:', error)
  }
}

const createNote = async () => {
  if (!newNote.value.trim()) return
  try {
    await axios.post(`${API_BASE_URL}/notes`, { content: newNote.value })
    newNote.value = ''
    await fetchNotes()
  } catch (error) {
    console.error('Error creating note:', error)
  }
}

const deleteNote = async (id) => {
  try {
    await axios.delete(`${API_BASE_URL}/notes/${id}`)
    await fetchNotes()
  } catch (error) {
    console.error('Error deleting note:', error)
  }
}

onMounted(fetchNotes)
</script>

<style>
.container {
  max-width: 600px;
  margin: auto;
  font-family: sans-serif;
  padding: 20px;
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
  background-color: #007bff;
  color: white;
  border: none;
  cursor: pointer;
}
button:hover {
  background-color: #0056b3;
}
li {
  list-style: none;
  margin-bottom: 10px;
  border-bottom: 1px solid #ccc;
  padding-bottom: 5px;
}
</style>
