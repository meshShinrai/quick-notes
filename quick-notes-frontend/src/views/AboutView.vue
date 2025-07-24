<template>
  <div class="about">
    <div class="content">
      <h1>About Quick Notes</h1>
      <p>
        Quick Notes is a simple and fast note-taking app that lets you jot down ideas, reminders, or to-dos in real-time.
        It's built with Vue 3 for the frontend and connects to a backend API for storing your notes.
      </p>

      <h2>How to Use</h2>
      <ul>
        <li>Go to the Home page</li>
        <li>Type your note in the input field</li>
        <li>Click "Add Note" to save it</li>
        <li>Click "Delete" next to a note to remove it</li>
      </ul>

      <p>Your notes are saved via the backend and reloaded every time you visit the app.</p>
    </div>
  </div>
</template>

<style>
.about {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: sans-serif;
  padding: 20px;
  background-color: #f8f9fa;
}
.content {
  max-width: 700px;
  background: white;
  padding: 30px;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}
h1 {
  margin-bottom: 10px;
}
h2 {
  margin-top: 20px;
}
ul {
  padding-left: 20px;
}
li {
  margin-bottom: 8px;
}
</style>
