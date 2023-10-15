<script setup>
import { defineProps, defineEmits, computed } from 'vue';

const props = defineProps({
    active: {
        type: Boolean,
    },
    info: {
        type: Object,
    }
})

const kilometers = computed(() => {
    if (props.info.distance) {
        const kilometers = props.info.distance / 1000;
        return kilometers >= 10 ? kilometers.toFixed(1) + ' км.': kilometers + ' км.';
    } else {
        return '';
    }
});
const emit = defineEmits(['closeFilter'])

const closeFilter = () => {
    emit('closeFilter')
}



function formatOpenHours(openHours) {
  if (!Array.isArray(openHours) || openHours.length === 0) {
    return 'Выходной';
  }

  const days = ['пн', 'вт', 'ср', 'чт', 'пт', 'сб', 'вс'];
  let currentDay = null;
  const formattedHours = [];

  for (const day of days) {
    const matchingDay = openHours.find(item => item.day === day);

    if (matchingDay) {
      if (matchingDay.hours === currentDay) {
        formattedHours[formattedHours.length - 1] += `, ${day}`;
      } else {
        formattedHours.push(`${day}: ${matchingDay.hours}`);
        currentDay = matchingDay.hours;
      }
    } else {
      formattedHours.push(`${day}: выходной`);
    }
  }

  return formattedHours.join(' ');
}
function generateRouteLink() {
    const myLatitude = 55.12345; 
    const myLongitude = 37.67890; 

    if (props.info && Object.keys(props.info).length > 0) {
        const destinationLatitude = props.info.position.latitude;
        const destinationLongitude = props.info.position.longitude;
        return `https://yandex.ru/maps/?rtext=С${myLatitude},${myLongitude}~${destinationLatitude},${destinationLongitude}&rtt=mt`;
    } else {
        return `https://yandex.ru/maps/?rtext=С${myLatitude},${myLongitude}~${myLatitude},${myLongitude}&rtt=mt`;
    }
}


</script>

<template>
    <div class="info"
        :class="{active: props.active}"
    >
        <div class="filter__header">
            <div class="filter__header-text">
                <div class="filter__header-title">
                    {{ props.info.sale_point_name }}
                </div>
            </div>
            <div class="filter__header-close"
                @click="closeFilter()"
            >
                <img src="@/assets/images/icons/close.svg">
            </div>
        </div>
        <div class="info__list custom-scroll">
            <div class="info__list-top">
                <p>
                    {{ props.info.address }}
                </p>
                <span>
                    {{ kilometers }}
                </span>
            </div>
            <div class="info__list-who">
                <div class="info__who-item">
                    <img class="icon" src="@/assets/images/icons/user.svg">
                    <span>
                        Физические лица
                    </span>
                </div>
                <div class="info__who-item">
                    <img class="icon" src="@/assets/images/icons/wheelchair.svg">
                    <span>
                        Маломобильные граждане
                    </span>
                </div>
                <div class="info__who-item">
                    <img class="icon" src="@/assets/images/icons/user-large.svg">
                    <span>
                        Клиенты с пакетом «Привилегия»
                    </span>
                </div>
                <div class="info__who-item">
                    <img class="icon" src="@/assets/images/icons/user-group.svg">
                    <span>
                        Организации
                    </span>
                </div>
            </div>
            <p class="info__subtitle" v-if="props.info.open_hours">
                Режим работы
            </p>
            <div class="info__time" v-if="props.info.open_hours">
                <p class="info__title">
                    Для физических лиц
                </p>
                <p class="time">
                    {{ formatOpenHours(props.info.open_hours) }}
                </p>
            </div>
            <div class="info__time" v-if="props.info.open_hours_individual">
                <p class="info__title">
                    Для организаций
                </p>
                <p class="time">
                    {{ formatOpenHours(props.info.open_hours_individual) }}
                </p>
            </div>
            <a :href="generateRouteLink()" target="_blank" class="info__btn btn btn_blue">
                Проложить маршрут на карте
            </a>
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
    padding: 20px 0px 0px;
    background: #1E1E1E;
    top: 100%;
    transition: .4s;
    overflow: auto;
    &.active {
        top: 0;
    }
    & .filter__header {
        border-bottom: 1px solid #2F3441;
        padding-bottom: 30px;
        margin-bottom: 0px;
    }
}

.info {
    display: flex;
    flex-direction: column;
&__list {
    padding: 20px 15px;
    flex: 1;
    overflow: auto;
}

&__list-top {
    display: flex;
    gap: 20px;
    justify-content: space-between;
    margin-bottom: 30px;
    & p {
        color: var(--absolute-input-gray, #ACB6C3);
        font-family: VTB Group UI;
        font-size: 18px;
        font-style: normal;
        font-weight: 400;
        line-height: 140%; /* 25.2px */
    }
    & span {
        color: var(--absolute-100, #FFF);
        font-family: VTB Group UI;
        font-size: 18px;
        font-style: normal;
        font-weight: 400;
        line-height: 140%; /* 25.2px */
        white-space: nowrap;
    }
}

&__list-who {
    display: flex;
    flex-wrap: wrap;
    gap: 20px;
    justify-content: space-between;
    margin-bottom: 40px;
}

&__who-item {
    flex: 0 0 calc(50% - 10px);
    color: var(--absolute-100, #FFF);
    font-family: VTB Group UI;
    font-size: 12px;
    font-style: normal;
    font-weight: 400;
    line-height: 120%; /* 14.4px */
    display: flex;
    gap: 8px;
    align-items: center;
}

&__subtitle {
    color: var(--absolute-100, #FFF);
    font-family: VTB Group UI;
    font-size: 16px;
    font-style: normal;
    font-weight: 400;
    line-height: 139%; /* 22.24px */
    margin-bottom: 25px;
}
&__time {
    margin-bottom: 30px;
}
&__btn {
    margin: 0 auto;
    width: fit-content;
}
& .time {
    color: var(--absolute-input-gray, #ACB6C3);
    font-family: VTB Group UI;
    font-size: 16px;
    font-style: normal;
    font-weight: 400;
    line-height: 139%; /* 22.24px */
}

&__title {
    color: var(--absolute-100, #FFF);
    font-family: VTB Group UI;
    font-size: 18px;
    font-style: normal;
    font-weight: 400;
    line-height: 140%; /* 25.2px */
    margin-bottom: 12px;
}
}

</style>