<template>
  <div>
    <h1>Notes App</h1>

    <form @submit.prevent="addNote">
      <input v-model="newNote.title" placeholder="Note Title" required />
      <textarea v-model="newNote.content" placeholder="Note Content" required></textarea>
      <button type="submit">Add Note</button>
    </form>

    <div v-if="notes.length">
      <h2>All Notes</h2>
      <ul>
        <li v-for="note in notes" :key="note.id">
          <strong>{{ note.title }}:</strong> {{ note.content }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const notes = ref([])
const newNote = ref({ title: '', content: '' })

const backendURL = import.meta.env.VITE_API_BASE_URL

const fetchNotes = async () => {
  try {
    const response = await axios.get(`${backendURL}/notes`)
    notes.value = response.data
  } catch (error) {
    console.error('Error fetching notes:', error)
  }
}

const addNote = async () => {
  try {
    await axios.post(`${backendURL}/notes`, newNote.value)
    newNote.value = { title: '', content: '' }
    fetchNotes()
  } catch (error) {
    console.error('Error adding note:', error)
  }
}

onMounted(fetchNotes)
</script>

<style scoped>
form {
  margin-bottom: 20px;
}
input,
textarea {
  display: block;
  margin-bottom: 10px;
  width: 100%;
}
button {
  padding: 8px 12px;
}
</style>
