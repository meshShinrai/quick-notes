<template>
  <div class="notes-app">
    <h1>Quick Notes</h1>

    <form @submit.prevent="addNote" class="note-form">
      <input v-model="newNote" type="text" placeholder="Enter a note..." required />
      <button type="submit">Add Note</button>
    </form>

    <div v-if="loading">Loading notes...</div>

    <ul class="note-list">
      <li v-for="note in notes" :key="note.id" class="note-item">
        <span>{{ note.content }}</span>
        <button @click="deleteNote(note.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'QuickNotes',
  data() {
    return {
      notes: [],
      newNote: '',
      loading: false,
    };
  },
  mounted() {
    this.fetchNotes();
  },
  methods: {
    async fetchNotes() {
      this.loading = true;
      try {
        const response = await fetch('https://your-eb-api-url/notes'); // <-- Replace with your real backend URL
        this.notes = await response.json();
      } catch (error) {
        console.error('Error fetching notes:', error);
      } finally {
        this.loading = false;
      }
    },
    async addNote() {
      try {
        const response = await fetch('https://your-eb-api-url/notes', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ content: this.newNote }),
        });
        const newNote = await response.json();
        this.notes.push(newNote);
        this.newNote = '';
      } catch (error) {
        console.error('Error adding note:', error);
      }
    },
    async deleteNote(id) {
      try {
        await fetch(`https://your-eb-api-url/notes/${id}`, { method: 'DELETE' });
        this.notes = this.notes.filter(note => note.id !== id);
      } catch (error) {
        console.error('Error deleting note:', error);
      }
    },
  },
};
</script>

<style scoped>
.notes-app {
  max-width: 600px;
  margin: 2rem auto;
  font-family: Arial, sans-serif;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
}

.note-form {
  display: flex;
  gap: 1rem;
  margin-bottom: 1rem;
}

.note-form input {
  flex: 1;
  padding: 0.5rem;
}

.note-form button {
  padding: 0.5rem 1rem;
}

.note-list {
  list-style: none;
  padding: 0;
}

.note-item {
  display: flex;
  justify-content: space-between;
  background: #f9f9f9;
  margin-bottom: 0.5rem;
  padding: 0.5rem;
  border-radius: 4px;
}
</style>
