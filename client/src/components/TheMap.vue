<script setup>
import * as ymaps3 from 'ymaps3';
import markImage from '@/assets/images/icons/mark.svg';
import markImage_purple from '@/assets/images/icons/mark_purple.svg';
import { ref, defineProps, defineEmits, watch } from 'vue';
import { useFilterStore } from '@/stores/filterStore';
let filterStore = useFilterStore();


const emit = defineEmits(['openInfoDepartament', 'mapInit'])

const props = defineProps({
    offices: {
        type: Array,
    }
})

const userCoordinates = ref(null)

async function getUserCoordinates() {
    return new Promise((resolve) => {
        // для демонстрации
        userCoordinates.value = [37.94, 55.76]
        resolve();
        // navigator.geolocation.getCurrentPosition(
        //     (position) => {
        //         const latitude = position.coords.latitude;
        //         const longitude = position.coords.longitude;
        //         userCoordinates.value = [longitude, latitude];
        //         resolve();
        //     },
        //     (error) => {
        //         console.error('Ошибка получения координат пользователя:', error);
        //         reject(error);
        //     }
        // );
    });
}
function circle(count) {
    const circle = document.createElement('div');
    circle.classList.add('circle');
    circle.innerHTML = `
            <div class="circle-content">
                <span class="circle-text">${count}</span>
            </div>
        `;
    return circle;
}

// Вызовите функцию, чтобы начать получать координаты пользователя
let map = null
let YMap = null;
let YMapDefaultSchemeLayer = null;
let YMapDefaultFeaturesLayer = null;
let YMapMarker = null;
let YMapFeatureDataSource = null;
let YMapLayer = null;
let YMapClusterer = null;
let clusterByGrid = null;
initMap();
async function initMap() {
    try {
        await getUserCoordinates();
        await ymaps3.ready;

        YMap = ymaps3.YMap;
        YMapDefaultSchemeLayer = ymaps3.YMapDefaultSchemeLayer;
        YMapDefaultFeaturesLayer = ymaps3.YMapDefaultFeaturesLayer;
        YMapMarker = ymaps3.YMapMarker;
        YMapFeatureDataSource = ymaps3.YMapFeatureDataSource;
        YMapLayer = ymaps3.YMapLayer;

        YMapClusterer = await ymaps3.import('@yandex/ymaps3-clusterer@0.0.1');
        clusterByGrid = YMapClusterer.clusterByGrid;
        YMapClusterer = YMapClusterer.YMapClusterer;


        // Координаты центра карты
        const CENTER_COORDINATES = userCoordinates.value;
        const LOCATION = {center: CENTER_COORDINATES, zoom: 9};

        const defaultSchemeLayer = new YMapDefaultSchemeLayer({theme: 'dark'});

        map = new YMap(document.getElementById('map'), {location: LOCATION});
        map.addChild(new YMapDefaultSchemeLayer())
        .addChild(new YMapDefaultFeaturesLayer())
        .addChild(new YMapFeatureDataSource({id: 'my-source'}))
        .addChild(new YMapLayer({source: 'my-source', type: 'markers', zIndex: 1800}))
        .addChild(defaultSchemeLayer);



        const contentPin = document.createElement('div');
        contentPin.classList.add('mark')
        contentPin.innerHTML = `<img src=${markImage}>`;


        filterStore.filter.position = {
            'latitude': userCoordinates.value[1],
            'longitude': userCoordinates.value[0]
        }
        // user
        const content = document.createElement('section');
        content.innerHTML = `<div class="me"></div>`;
        const userMark = new YMapMarker({       
            coordinates: userCoordinates.value,
            draggable: false
        },
            content
        );
        map.addChild(userMark);

        emit('mapInit')
    } catch (err) {
        console.log(err)
    }
}

let clusterer = null;
const createCluster = (marks) => {
    if(!marks) return
    console.log(map.children)

    map.destroy()
    const defaultSchemeLayer = new YMapDefaultSchemeLayer({theme: 'dark'});
    const CENTER_COORDINATES = userCoordinates.value;
    const LOCATION = {center: CENTER_COORDINATES, zoom: 9};
    map = new YMap(document.getElementById('map'), {location: LOCATION});
        map.addChild(new YMapDefaultSchemeLayer())
        .addChild(new YMapDefaultFeaturesLayer())
        .addChild(new YMapFeatureDataSource({id: 'my-source'}))
        .addChild(new YMapLayer({source: 'my-source', type: 'markers', zIndex: 1800}))
        .addChild(defaultSchemeLayer);
            // user
    const content = document.createElement('section');
        content.innerHTML = `<div class="me"></div>`;
        const userMark = new YMapMarker({       
            coordinates: userCoordinates.value,
            draggable: false
        },
            content
    );
    map.addChild(userMark);
    const marker = (feature) => {
        const contentPin = document.createElement('div');
        contentPin.classList.add('mark');
        contentPin.classList.add('load');
        if(feature.load <= 20 && feature.load >= 0) {
            contentPin.classList.add(`load_low`);
        }
        if(feature.load <= 40 && feature.load > 20) {
            contentPin.classList.add(`load_mid`);
        }
        if(feature.load > 40) {
            contentPin.classList.add(`load_hard`);
        }

        contentPin.setAttribute('data-id', feature.id);
        if(filterStore.filter.officeType === 'Да') {
            contentPin.innerHTML = `<img src=${markImage}>`;
        } 
        if(filterStore.filter.officeType === 'Да (Зона Привилегия)') {
            contentPin.innerHTML = `<img src=${markImage_purple}>`;
        }

        const markerElement = new ymaps3.YMapMarker({
            coordinates: feature.geometry.coordinates,
            source: 'my-source',
            onClick: () => {
                emit('openInfoDepartament', feature.id);
            }
        }, contentPin.cloneNode(true));

        return markerElement;
    };

    const cluster = (coordinates, features) =>
        new ymaps3.YMapMarker({
            coordinates,
            source: 'my-source'
        },
        circle(features.length).cloneNode(true)
    );


    const points = marks.map((mark) => ({
        type: 'Feature',
        id: mark.id,
        load: mark.load,
        geometry: { coordinates: [mark.position.longitude, mark.position.latitude] },
        properties: { name: 'Point of issue of orders'}
    }));

    clusterer = new YMapClusterer({
        method: clusterByGrid({ gridSize: 64 }),
        features: points,
        marker,
        cluster
    });

    map.addChild(clusterer);
}

watch(() => props.offices,
    () => {
        createCluster(props.offices)
    },
    {deep: true}
)
</script>

<template>
    <div id="map"></div>
</template>


<style lang="scss">
#map {
    width: 100%;
    height: calc(100vh - 71px);
    @media (max-width: 539px) {
        height: calc(50vh - 71px);
    }
}
.me {
    width: 18px;
    height: 18px;
    box-sizing: border-box;
    background-color: rgb(13, 105, 242);
    border-radius: 50%;
    border: 2px solid rgb(255, 255, 255);
    -webkit-box-shadow: 0px 0px 20px 20px rgba(34, 60, 80, 0.4);
    -moz-box-shadow: 0px 0px 20px 20px rgba(34, 60, 80, 0.2);
    box-shadow: 0px 0px 20px 20px rgba(34, 60, 80, 0.2);
    transform: translate(-50%, -50%);
}

.circle {
    background: #fff;
    border: 2px solid blue;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 35px;
    height: 35px;
    border-radius: 50%;
}

.mark {
    cursor: pointer;
    transform: translate(-14px, -42px);
    &::after {
        content: "";
        width: 5px;
        height: 5px;
        background: red;
        border-radius: 50%;
        position: absolute;
        left: 0;
        top: 0;
    }
}
.ymaps3x0--marker {
    position: relative;
}
.load {
    &::after {
        content: "";
        width: 7px;
        height: 7px;
        border-radius: 50%;
        position: absolute;
        right: 0;
        top: 0;
        background: transparent;
    }
    &_low {
        &::after {
            background: green;
        }
    }
    &_mid {
        &::after {
            background: orange;
        }
    }
    &_hard {
        &::after {
            background: red;
        }
    }
}
</style>