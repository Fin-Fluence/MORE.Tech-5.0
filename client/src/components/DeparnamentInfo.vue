<script setup>
import { defineProps, defineEmits } from 'vue';

const props = defineProps({
    active: {
        type: Boolean,
    }
})

const emit = defineEmits(['closeFilter'])

const closeFilter = () => {
    emit('closeFilter')
}
</script>

<template>
    <div class="info"
        :class="{active: props.active}"
    >
    <div class="filter__header">
        <div class="filter__header-text">
            <div class="filter__header-title">
                Данные отделения
            </div>
        </div>
        <div class="filter__header-close"
            @click="closeFilter()"
        >
            <img src="@/assets/images/icons/close.svg">
        </div>
    </div>
    </div>
</template>

<style lang="scss" scoped>
.info {
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    z-index: 2;
    padding: 20px 0px;
    background: #1E1E1E;
    top: 100%;
    transition: .4s;
    overflow: auto;
    &.active {
        top: 0;
    }
}
</style>