<script setup>
import TheDepartament from './TheDepartament.vue';

import { defineEmits } from 'vue';

const emit = defineEmits(['openInfoDepartament'])

const openInfoDepartament = () => {
    emit('openInfoDepartament')
}
</script>

<template>
    <div class="departaments">
        <div class="departaments__subtitle">
          Отделения в радиусе 3 км:
        </div>
        <the-departament
            v-for="(departament, index) in [1,2,3]" :key="index"
            :workLoad="index"
            @openInfoDepartament="openInfoDepartament()"
        />
    </div>
</template>

<style lang="scss" scoped>
.departaments {
    padding-right: 2px;
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