<template>
  <div>
    <form @submit.prevent="addNote">
      <input v-model="newNote" placeholder="Write a note..." />
      <button type="submit">Add Note</button>
    </form>

    <ul>
      <li v-for="note in notes" :key="note.id">
        {{ note.content }}
        <button @click="deleteNote(note.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const notes = ref([])
const newNote = ref('')

const API = 'http://localhost:5000/notes'

const fetchNotes = async () => {
  try {
    const res = await axios.get(API)
    notes.value = res.data
  } catch (err) {
    console.error('Error fetching notes:', err)
  }
}

const addNote = async () => {
  if (!newNote.value.trim()) return
  try {
    const res = await axios.post(API, { content: newNote.value })
    notes.value.push(res.data)
    newNote.value = ''
  } catch (err) {
    console.error('Error adding note:', err)
  }
}

const deleteNote = async (id) => {
  try {
    await axios.delete(`${API}/${id}`)
    notes.value = notes.value.filter(note => note.id !== id)
  } catch (err) {
    console.error('Error deleting note:', err)
  }
}

onMounted(fetchNotes)
</script>

<style scoped>
form {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

input {
  flex: 1;
  padding: 0.5rem;
}

button {
  padding: 0.5rem;
  cursor: pointer;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  margin-bottom: 0.5rem;
  display: flex;
  justify-content: space-between;
  background: #f9f9f9;
  padding: 0.5rem;
}
</style>
