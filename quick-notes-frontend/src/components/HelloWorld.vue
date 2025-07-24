<template>
  <div>
    <h1>Quick Notes</h1>

    <form @submit.prevent="addNote">
      <textarea v-model="newNote" placeholder="Write your note here..." required></textarea>
      <button type="submit">Add Note</button>
    </form>

    <div v-if="notes.length">
      <h2>Your Notes</h2>
      <ul>
        <li v-for="note in notes" :key="note.id">
          <p>{{ note.content }}</p>
          <small>{{ new Date(note.created_at).toLocaleString() }}</small>
          <button @click="deleteNote(note.id)">Delete</button>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const notes = ref([])
const newNote = ref('')

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

const fetchNotes = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/api/notes`)
    notes.value = response.data
  } catch (error) {
    console.error('Error fetching notes:', error)
  }
}

const addNote = async () => {
  if (!newNote.value.trim()) return
  
  try {
    await axios.post(`${API_BASE_URL}/api/notes`, { 
      content: newNote.value 
    })
    newNote.value = ''
    await fetchNotes() // Refresh the list
  } catch (error) {
    console.error('Error adding note:', error)
  }
}

const deleteNote = async (id) => {
  try {
    await axios.delete(`${API_BASE_URL}/api/notes/${id}`)
    await fetchNotes() // Refresh the list
  } catch (error) {
    console.error('Error deleting note:', error)
  }
}

onMounted(fetchNotes)
</script>

<style scoped>
form {
  margin-bottom: 2rem;
  max-width: 500px;
}

textarea {
  width: 100%;
  padding: 0.5rem;
  margin-bottom: 1rem;
  min-height: 100px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  background: #42b983;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background: #3aa876;
}

ul {
  list-style: none;
  padding: 0;
  max-width: 600px;
}

li {
  background: #f9f9f9;
  padding: 1rem;
  margin-bottom: 1rem;
  border-radius: 4px;
  display: flex;
  flex-direction: column;
}

li p {
  margin: 0 0 0.5rem 0;
}

li small {
  color: #666;
  font-size: 0.8rem;
  margin-bottom: 0.5rem;
}

li button {
  align-self: flex-end;
  background: #ff4444;
}

li button:hover {
  background: #cc0000;
}
</style>
