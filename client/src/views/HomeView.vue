<script setup>
import TheFilter from '@/components/TheFilter.vue';
import DeparnamentInfo from '@/components/DeparnamentInfo.vue';
import TheMap from '@/components/TheMap.vue';
import { onMounted, ref } from 'vue';
import { useFilterStore } from '@/stores/filterStore';
import TheDepartament from '@/components/TheDepartament.vue';

const currentDeparnamentInfo = ref({})

const filterCurrentFepartament = (id) => {
  currentDeparnamentInfo.value = {...offices.value.find(department => department.id === id)}
};


const toggleBtn = (value) => {
    console.log(filterStore.filter.service_names)
    if (filterStore.filter.service_names.includes(value)) {
        filterStore.removeServiceName(value);
    } else {
        filterStore.addServiceName(value);
    }
    getOffice()
}

let filterStore = useFilterStore();

let filterIsActive = ref(false)
let infoIsActive = ref(false)
let openFullContent = ref(false)

const openInfoDepartament = () => {
  infoIsActive.value = true
}

const openContentFull = () => {
  openFullContent.value = !openFullContent.value
}

const paramsObjects = ref('office')

const offices = ref([]);
const getOffice = async () => {
  try {
    const params = { ...filterStore.filter };
    if (Array.isArray(params.service_names) && params.service_names.length === 0) {
      delete params.service_names;
    }
    const url = `http://localhost:3000/${paramsObjects.value}?filter=` + JSON.stringify(params);
    console.log(url)
    const response = await fetch(url, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });
    const data = await response.json();
    offices.value = data;
  } catch (err) {
    console.error('Произошла ошибка:', err);
  }
};

const getNewAllObjects = () => {
  getOffice()
}

const mapInit = () => {
  getOffice();
}

onMounted(() => {
})
</script>

<template>
  <div class="home">
    <div class="content"
      :class="{openFull: openFullContent}"
    >
      <div class="mobile-btn"
        @click="openContentFull"
      >

      </div>
      <the-filter
        :active="filterIsActive"
        :activeObject="paramsObjects"
        @closeFilter="() => filterIsActive = false"
        @getMarksWithFilter="getOffice()"
      />
      <deparnament-info
        :active="infoIsActive"
        @closeFilter="() => infoIsActive = false"
        :info="currentDeparnamentInfo"
      />
      <div class="content__header">
        <div class="content__search">
          <input placeholder="Город, район, метро, улица">
          <img class="icon search" src="@/assets/images/icons/search.svg">
          <img class="icon close" src="@/assets/images/icons/close.svg">
        </div>
        <div class="search-main">
          <div class="search-main__top">
            <div class="search-main__top-btns">
              <div class="search-main__top-recent">
              Недавние
              </div>
              <button class="search-main__top-reset">
                Сбросить
              </button>
            </div>
            <img class="close" src="@/assets/images/icons/close.svg">
          </div>
        </div>
        <div class="content__filter">
          <div class="content__filter-top">
            <div class="content__filter-btns">
              <button class="btn "
                :class="{btn_blue: paramsObjects === 'office'}"
                @click="() => {paramsObjects = 'office', getNewAllObjects()}"
              >
                Отделения
              </button>
              <button class="btn"
                :class="{btn_blue: paramsObjects === 'atm'}"
                @click="() => {paramsObjects = 'atm', getNewAllObjects()}"
              >
                Банкоматы
              </button>
            </div>
            <button class="btn btn_settings"
              @click="() => filterIsActive = true"
            >
              <svg xmlns="http://www.w3.org/2000/svg" width="21" height="14" viewBox="0 0 21 14" fill="none">
                <path d="M11.3867 9.00004C11.8672 7.82087 12.9688 7.00004 14.25 7.00004C15.5312 7.00004 16.6328 7.82087 17.1133 9.00004H19.25C19.9414 9.00004 20.5 9.59587 20.5 10.3334C20.5 11.0709 19.9414 11.6667 19.25 11.6667H17.1133C16.6328 12.8459 15.5312 13.6667 14.25 13.6667C12.9688 13.6667 11.8672 12.8459 11.3867 11.6667H1.75C1.05859 11.6667 0.5 11.0709 0.5 10.3334C0.5 9.59587 1.05859 9.00004 1.75 9.00004H11.3867Z" fill="white"/>
                <path d="M8 0.333374C9.28125 0.333374 10.3828 1.15421 10.8633 2.33337H19.25C19.9414 2.33337 20.5 2.92921 20.5 3.66671C20.5 4.40421 19.9414 5.00004 19.25 5.00004H10.8633C10.3828 6.17921 9.28125 7.00004 8 7.00004C6.71875 7.00004 5.61719 6.17921 5.13672 5.00004H1.75C1.05859 5.00004 0.5 4.40421 0.5 3.66671C0.5 2.92921 1.05859 2.33337 1.75 2.33337H5.13672C5.61719 1.15421 6.71875 0.333374 8 0.333374Z" fill="white"/>
              </svg>
            </button>
          </div>
          <div class="content__filter-whom" v-if="paramsObjects === 'office'">
            <div class="subtitle">
              Выбор сегмента
            </div>
            <div class="content__whom-list">
              <button class="content__whom-btn btn btn_blue">
                Физическим лицам
              </button>
              <button class="content__whom-btn btn">
                Организациям
              </button>
              <button class="content__whom-btn btn">
                Прайм
              </button>
              <button class="content__whom-btn btn">
                Привелигированным
              </button>
            </div>
          </div>
          <div class="content__filter-whom" v-if="paramsObjects === 'office'">
            <div class="subtitle">
              Быстрый доступ
            </div>
            <div class="content__whom-list">
              <button class="content__whom-btn btn"
                :class="{btn_blue: filterStore.filter.service_names.includes('BrokerService')}"
                @click="toggleBtn('BrokerService')"
              >
                Конвертация валют
              </button>
              <button class="content__whom-btn btn"
                :class="{btn_blue: filterStore.filter.service_names.includes('ConversionFL')}"
                @click="toggleBtn('ConversionFL')"
              >
                Брокерское обслуживание
              </button>
            </div>
          </div>
        </div>
      </div>
      <div class="departaments custom-scroll">
        <div class="departaments">
          <the-departament
              v-for="(departament, index) in offices" :key="index"
              :departament="departament"
              @openInfoDepartament="openInfoDepartament()"
              @click="filterCurrentFepartament(departament.id)"
          />
      </div>
      </div>
    </div>
    <the-map
      :offices="offices"
      @openInfoDepartament="(id) => {filterCurrentFepartament(id), openInfoDepartament()}"
      @mapInit="mapInit()"
    />
  </div>
</template>

<style lang="scss">
.mobile-btn {
  display: none;
  width: 70px;
  height: 4px;
  background: rgb(226, 229, 233);
  position: absolute;
  top: 5px;
  border-radius: 12px;
  left: 50%;
  transform: translate(-50%, 0);
  @media (max-width: 539px) {
    display: block;
  }
}
.departaments {
  overflow: auto;
  flex: 1;
}
.home {
  position: relative;
  display: flex;
  @media (max-width: 539px) {
    flex-direction: column-reverse;
  }
}
.content {
  background: #1E1E1E;
  padding: 40px 16px 0px;
  width: auto;
  max-width: 425px;
  height: 100%;
  display: flex;
  flex-direction: column;
  height: calc(100vh - 71px);
  position: relative;
  overflow: hidden;
  @media (max-width: 539px) {
    transition: .2s;
    position: absolute;
    z-index: 5;
    height: calc(100vh - 71px);
    top: 100%;
    width: 100%;
    max-width: none;
    &.openFull {
      top: 0px;
    }
  }
  &__header {

  }

  &__search {
    background: #2f3441;
    position: relative;
    border-radius: 8px;
    margin-bottom: 30px;
    & input {
      width: 100%;
      background: transparent;
      padding: 0 27px;
      padding: 14px 33px;
    }
    & .icon {
      position: absolute;
      top: 50%;
      transform: translate(0, -50%);
    }
    & .close {
      right: 10px;
      opacity: 0;
      cursor: pointer;
      
    }
    & .search {
      left: 10px;
    }
  }

  &__filter {

  }

  &__filter-btns {
    border-radius: 8px;
    background: #2f3441;
    display: flex;
    flex: 1;
    & .btn {
      flex: 1;
    }
  }

  &__filter-top {
    display: flex;
    gap: 18px;
    margin-bottom: 40px;
  }

  &__filter-whom {
    margin-bottom: 40px;

  }

  &__whom-list {
    display: flex;
    flex-wrap: wrap;
    gap: 12px;
    @media (max-width: 539px) {
      flex-wrap: nowrap;
      overflow: auto;
    }
  }

  &__whom-btn {
    border: 1px solid #4789EB;
    @media (max-width: 539px) {
      white-space: nowrap;
    }
    &.btn {
      color: #fff;
      padding: 7px 20px;
      background: transparent;
      border-radius: 30px;
      font-family: 'Inter';
      font-size: 14px;
      font-style: normal;
      font-weight: 400;
      line-height: 140%; /* 19.6px */
    }
    &.btn_blue {
      border: 1px solid transparent;
      background: var(--gradient, linear-gradient(90deg, #0037FF 0%, #0085FF 100%));
    }
  }
}
.search-main {
  background: #2f3441;
  padding: 20px 10px 21px 10px;
  border-radius: 12px;
  position: absolute;
  width: 100%;
  left: 0;
  top: 100px;
  display: none;
&__top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 10px;

}
&__top-btns {
  display: flex;
  gap: 20px;
  align-items: flex-end;
}
&__top-recent {
  color: var(--absolute-stroke-gray, #DCE0EB);
  font-family: Inter;
  font-size: 16px;
  font-style: normal;
  font-weight: 600;
  line-height: 120%; /* 19.2px */
}

&__top-reset {
  color: var(--absolute-input-gray, #ACB6C3);
  font-family: Inter;
  font-size: 12px;
  font-style: normal;
  font-weight: 400;
  line-height: 120%; /* 14.4px */
}
}
.close {
  cursor: pointer;
}

</style>