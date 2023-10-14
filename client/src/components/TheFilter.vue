
<script setup>

</script>

<template>
    <div class="filter">
        <div class="filter__header">
            <div class="filter__header-text">
                <div class="filter__header-title">
                    filter
                </div>
                <div class="filter__header-title">
                    Сбросить
                </div>
            </div>
            <div class="filter__header-close">
                <img src="@/assets/images/icons/close.svg">
            </div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.filter {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    z-index: 2;
    padding: 20px;
    background: #1E1E1E;
    display: none;
    &__header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 30px;
        width: 100%;
    }

    &__header-text {
        display: flex;
        align-items: flex-end;
        gap: 30px;
    }

    &__header-title {

    }

    &__header-close {

    }

}
</style>