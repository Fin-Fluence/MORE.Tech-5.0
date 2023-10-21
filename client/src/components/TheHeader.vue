<script setup>

</script>

<template>
    <header class="header-w">
        <div class="header container">
            <img src="@/assets/images/icons/logo.svg">
            <img class="header__robot" src="@/assets/images/icons/robot.svg">
        </div>
    </header>
</template>

<style lang="scss" scoped>
.header {
    display: flex;
    align-items: center;
    justify-content: space-between;
}
.header-w {
    background: #2B313B;
    padding: 10px 0px;
}
.header__robot {
  width: 70px;
}
</style>