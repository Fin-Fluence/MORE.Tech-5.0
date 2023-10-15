<script setup>

</script>

<template>
    <header class="header-w">
        <div class="header container">
            <img src="@/assets/images/icons/logo.svg">
            <!-- <div class="header__menu ">
                <div class="nav-icon toggle-menu">
                    <span></span>
                    <span></span>
                    <span></span>
                </div>
            </div> -->
            <img src="@/assets/images/icons/robot.svg">
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
    padding: 20px 0px;
}
.nav-icon {
  width: 22px;
  height: 15px;
  position: relative;
  transform: rotate(0deg);
  transition: .5s ease-in-out;
  cursor: pointer;
}
.nav-icon span {
  display: block;
  position: absolute;
  height: 2.62px;
  width: 100%;
  background: #fff;
  border-radius: 10px;
  opacity: 1;
  left: 0;
  transform: rotate(0deg);
  transition: .25s ease-in-out;
}

.nav-icon span:nth-child(1) {
  top: 0px;
}
.nav-icon span:nth-child(2) {
  top: 6px;
}
.nav-icon span:nth-child(3) {
  top: 12px;
}
.nav-icon.open span {
  left: 7px;
}
.nav-icon.open span:nth-child(1) {
  top: 5px;
  transform: rotate(135deg);
  max-width: 30px;
  left: 0px;
}
.nav-icon.open span:nth-child(2) {
  opacity: 0;
  left: -60px;
}
.nav-icon.open span:nth-child(3) {
  top: 5px;
  transform: rotate(-135deg);
  max-width: 30px;
  left: 0px;
}
.header__menu {
  display: none;
  @media (max-width: 1023px) {
    display: block;
  }
}
</style>