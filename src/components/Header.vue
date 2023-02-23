<script setup lang="ts"></script>

<template>
  <header>
    <nav>
      <ul>
        <li><RouterLink to="/">Home</RouterLink></li>
        <li><a href="#news">Page1</a></li>
        <li><a href="#contact">Page2</a></li>
        <li style="float: right">
          <RouterLink to="/about" class="active">About</RouterLink>
        </li>
      </ul>
    </nav>
  </header>
</template>

<style scoped>
header {
  height: 45px;
  display: block;
  background-color: #fff;
  border-bottom: 1px solid #d2d2d2;
}
ul {
  list-style-type: none;
  margin: 0;
  padding: 0 0 0 10px;
  overflow: hidden;
}
li {
  display: inline;
  float: left;
}

li a {
  padding: 8px;
  display: block;
  color: black;
  text-align: center;
  padding: 12px 16px;
  text-decoration: none;
}

li a:hover {
  background-color: #d1d1d1;
}
li a.active {
  background-color: #d1d1d1;
}
</style>
