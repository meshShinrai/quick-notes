<template>
  <div class="notes-app">
    <h1>Quick Notes</h1>

    <form @submit.prevent="addNote" class="note-form">
      <textarea 
        v-model="newNote" 
        placeholder="Enter your note here..." 
        required
        rows="3"
      ></textarea>
      <button type="submit" class="submit-btn">Add Note</button>
    </form>

    <div v-if="loading" class="loading-state">
      <p>Loading notes...</p>
    </div>
    <div v-else-if="error" class="error-state">
      <p>{{ error }}</p>
    </div>

    <ul v-else class="note-list">
      <li v-for="note in notes" :key="note.id" class="note-item">
        <div class="note-content">
          <p>{{ note.content }}</p>
          <small class="note-date">{{ formatDate(note.created_at) }}</small>
        </div>
        <button 
          @click="deleteNote(note.id)" 
          class="delete-btn"
          aria-label="Delete note"
        >
          Delete
        </button>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'QuickNotes',
  data() {
    return {
      notes: [],
      newNote: '',
      loading: false,
      error: null,
    };
  },
  computed: {
    apiBaseUrl() {
      return import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';
    }
  },
  methods: {
    formatDate(dateString) {
      return new Date(dateString).toLocaleString();
    },
    async fetchNotes() {
      this.loading = true;
      this.error = null;
      try {
        const response = await axios.get(`${this.apiBaseUrl}/api/notes`);
        this.notes = response.data;
      } catch (error) {
        console.error('Error fetching notes:', error);
        this.error = 'Failed to load notes. Please try again.';
      } finally {
        this.loading = false;
      }
    },
    async addNote() {
      if (!this.newNote.trim()) return;
      
      this.error = null;
      try {
        const response = await axios.post(`${this.apiBaseUrl}/api/notes`, {
          content: this.newNote
        });
        this.notes.unshift(response.data); // Add new note at beginning
        this.newNote = '';
      } catch (error) {
        console.error('Error adding note:', error);
        this.error = 'Failed to add note. Please try again.';
      }
    },
    async deleteNote(id) {
      this.error = null;
      try {
        await axios.delete(`${this.apiBaseUrl}/api/notes/${id}`);
        this.notes = this.notes.filter(note => note.id !== id);
      } catch (error) {
        console.error('Error deleting note:', error);
        this.error = 'Failed to delete note. Please try again.';
      }
    },
  },
  mounted() {
    this.fetchNotes();
  },
};
</script>

<style scoped>
.notes-app {
  max-width: 600px;
  margin: 2rem auto;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  background: white;
}

h1 {
  color: #2c3e50;
  text-align: center;
  margin-bottom: 1.5rem;
}

.note-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-bottom: 2rem;
}

textarea {
  width: 100%;
  padding: 0.8rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
  resize: vertical;
  min-height: 80px;
}

.submit-btn {
  padding: 0.8rem;
  background-color: #42b983;
  color: white;
  border: none;
  border-radius: 4px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.2s;
}

.submit-btn:hover {
  background-color: #3aa876;
}

.loading-state,
.error-state {
  text-align: center;
  padding: 1rem;
  margin: 1rem 0;
  border-radius: 4px;
}

.loading-state {
  background-color: #f8f9fa;
  color: #495057;
}

.error-state {
  background-color: #f8d7da;
  color: #721c24;
}

.note-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.note-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  margin-bottom: 1rem;
  background-color: #f8f9fa;
  border-radius: 4px;
  transition: transform 0.2s;
}

.note-item:hover {
  transform: translateY(-2px);
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

.delete-btn {
  padding: 0.5rem 1rem;
  background-color: #dc3545;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.delete-btn:hover {
  background-color: #c82333;
}

@media (max-width: 640px) {
  .notes-app {
    margin: 1rem;
    padding: 1rem;
  }
}
</style>
