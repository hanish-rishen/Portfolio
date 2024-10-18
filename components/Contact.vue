<template>
  <section id="contact" class="min-h-screen flex flex-col justify-center items-center p-4">
    <h2 class="title text-3xl md:text-4xl mb-4 text-center" @mouseenter="expandCursor" @mouseleave="shrinkCursor">Contact</h2>
    <p class="subtitle text-lg md:text-xl text-center" @mouseenter="expandCursor" @mouseleave="shrinkCursor">Get in touch</p>
    <!-- Add your contact content here -->
  </section>
</template>

<script setup>
const expandCursor = () => {
  // Implement cursor expansion logic
}

const shrinkCursor = () => {
  // Implement cursor shrinking logic
}
</script>