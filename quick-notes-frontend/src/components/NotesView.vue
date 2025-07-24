<template>
  <div class="notes-container">
    <h1>Quick Notes</h1>
    
    <form @submit.prevent="addNote" class="note-form">
      <input
        v-model="newNote"
        placeholder="Write a note..."
        class="note-input"
      />
      <button type="submit" class="add-button">
        Add Note
      </button>
    </form>

    <div v-if="loading" class="loading">Loading notes...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    
    <ul v-else class="notes-list">
      <li v-for="note in notes" :key="note.id" class="note-item">
        <div class="note-content">
          <p>{{ note.content }}</p>
          <small class="note-date">
            {{ formatDate(note.created_at) }}
          </small>
        </div>
        <button
          @click="deleteNote(note.id)"
          class="delete-button"
          aria-label="Delete note"
        >
          Delete
        </button>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const notes = ref([])
const newNote = ref('')
const loading = ref(true)
const error = ref(null)

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'
const API_URL = `${API_BASE_URL}/api/notes`

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleString()
}

const fetchNotes = async () => {
  try {
    loading.value = true
    const response = await axios.get(API_URL)
    notes.value = response.data
  } catch (err) {
    console.error('Error fetching notes:', err)
    error.value = 'Failed to load notes. Please try again later.'
  } finally {
    loading.value = false
  }
}

const addNote = async () => {
  if (!newNote.value.trim()) return
  
  try {
    const response = await axios.post(API_URL, { 
      content: newNote.value 
    })
    notes.value.unshift(response.data) // Add new note at beginning
    newNote.value = ''
  } catch (err) {
    console.error('Error adding note:', err)
    error.value = 'Failed to add note. Please try again.'
  }
}

const deleteNote = async (id) => {
  try {
    await axios.delete(`${API_URL}/${id}`)
    notes.value = notes.value.filter(note => note.id !== id)
  } catch (err) {
    console.error('Error deleting note:', err)
    error.value = 'Failed to delete note. Please try again.'
  }
}

onMounted(fetchNotes)
</script>

<style scoped>
.notes-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 1rem;
}

h1 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 1.5rem;
}

.note-form {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
}

.note-input {
  flex: 1;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

.add-button {
  padding: 0.75rem 1.5rem;
  background-color: #42b983;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
  transition: background-color 0.2s;
}

.add-button:hover {
  background-color: #3aa876;
}

.notes-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.note-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  margin-bottom: 0.75rem;
  background-color: #f8f9fa;
  border-radius: 4px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.note-content {
  flex: 1;
}

.note-date {
  display: block;
  color: #6c757d;
  font-size: 0.8rem;
  margin-top: 0.5rem;
}

.delete-button {
  padding: 0.5rem 1rem;
  background-color: #dc3545;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.delete-button:hover {
  background-color: #c82333;
}

.loading,
.error {
  text-align: center;
  padding: 1rem;
  margin: 1rem 0;
  border-radius: 4px;
}

.loading {
  background-color: #e9ecef;
  color: #495057;
}

.error {
  background-color: #f8d7da;
  color: #721c24;
}
</style>
