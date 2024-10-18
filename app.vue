<template>
  <div class="bg-black text-white">
    <CustomCursor />
    <Navigation />
    <Header />
    <main>
      <About />
      <Projects />
      <Skills />
      <Contact />
    </main>
  </div>
</template>

<script setup>
import CustomCursor from './components/CustomCursor.vue'
import Navigation from './components/Navigation.vue'
import Header from './components/Header.vue'
import About from './components/About.vue'
import Projects from './components/Projects.vue'
import Skills from './components/Skills.vue'
import Contact from './components/Contact.vue'
</script>

<style>
body {
  font-family: 'Satoshi', sans-serif;
  background-color: #000;
  color: #fff;
  margin: 0;
  padding-top: 60px;
  cursor: none;
}

.title {
  font-family: 'Cabinet Grotesk', sans-serif;
  font-weight: 100;
}

.subtitle {
  font-style: italic;
}

a, button {
  cursor: none;
}
</style>
