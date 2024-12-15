<template>
  <div class="product-detail">
    <h1>Product Detail</h1>
    <p>Product ID: {{ $route.params.id }}</p>
  </div>
</template>

<script>
export default {
  name: 'ProductDetail',
};
</script>