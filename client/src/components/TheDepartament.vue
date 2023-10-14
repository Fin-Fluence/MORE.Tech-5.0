<script setup>
import { defineProps, defineEmits } from 'vue';

const emit = defineEmits(['openInfoDepartament'])

const props = defineProps({
    workLoad: {
        type: Boolean,
    }
})

const openInfoDepartament = () => {
    emit('openInfoDepartament')
}
</script>

<template>
    <div class="departament"
        @click="openInfoDepartament"
    >
        <div class="departament__icon">
            <img src="@/assets/images/icons/mark.svg">
        </div>
        <div class="departament__text">
            <div class="departament__address">
                г. Москва, Ленинский пр-т, д. 34/1
            </div>
            <div class="departament__info">
                <span v-if="props.workLoad === 0" style="color: #F1A038;">
                    заполнено
                </span>
                <span v-if="props.workLoad === 1">
                    не заполнено
                </span>
                <span v-if="props.workLoad === 2" style="color: #CA181F;">
                    переполнено
                </span>
                <div class="point">

                </div>
                <span>
                    закроется в 19:00
                </span>
            </div>
        </div>
        <div class="departament__distance">
            2.4 км
        </div>
    </div>
</template>

<style lang="scss" scoped>
.departament {
    display: flex;
    align-items: center;
    padding: 12px;
    cursor: pointer;
    border: 1px solid transparent;
    border-radius: 12px;
    &:hover {
        border: 1px solid #4789EB;
    }
    &__icon {
        margin-right: 16px;
    }

    &__text {
        margin-right: 6px;
    }

    &__address {
        color: var(--absolute-100, #FFF);
        font-family: VTB Group UI;
        font-size: 18px;
        font-style: normal;
        font-weight: 400;
        line-height: 140%;
        margin-bottom: 10px;
    }

    &__info {
        color: #FFF;
        font-family: Inter;
        font-size: 14px;
        font-style: normal;
        font-weight: 400;
        display: flex;
        align-items: center;
        gap: 5px;
        & .point {
            width: 5px;
            height: 5px;
            background: #fff;
            border-radius: 50%;
            margin-top: 3px;
        }
        & span {
            display: block;
        }
    }

    &__distance {
        color: #FFF;
        font-family: VTB Group UI;
        font-size: 18px;
        font-style: normal;
        font-weight: 400;
        line-height: 140%; /* 25.2px */
    }
}

</style>