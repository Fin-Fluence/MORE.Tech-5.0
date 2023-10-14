<script setup>
import TheDepartament from './TheDepartament.vue';

import { defineEmits, defineProps } from 'vue';

const emit = defineEmits(['openInfoDepartament'])

const openInfoDepartament = () => {
    emit('openInfoDepartament')
}

const props = defineProps({
    departaments: {
        type: Array,
        required: true
    },
    title: {
        type: String,
        required: true
    }
})
</script>

<template>
    <div class="departaments">
        <div class="departaments__subtitle">
            {{ props.title }}
        </div>
        <the-departament
            v-for="(departament, index) in [1,2]" :key="index"
            :departament="departament"
            @openInfoDepartament="openInfoDepartament()"
        />
    </div>
</template>

<style lang="scss" scoped>
.departaments {
    padding-right: 2px;
    padding-bottom: 10px;
    &__subtitle {
        color: var(--absolute-input-gray, #ACB6C3);
        font-family: Inter;
        font-size: 16px;
        font-style: normal;
        font-weight: 600;
        line-height: 120%; /* 19.2px */
        margin-bottom: 16px;
    }
}

</style>