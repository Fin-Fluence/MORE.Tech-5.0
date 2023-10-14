<script setup>

</script>

<template>
    <div class="departament">
        <div class="departament__icon">
            <img src="@/assets/images/icons/mark.svg">
        </div>
        <div class="departament__text">
            <div class="departament__address">
                г. Москва, Ленинский пр-т, д. 34/1
            </div>
            <div class="departament__info">
                <span>
                    не заполнено
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