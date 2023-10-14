<script setup>
import * as ymaps3 from 'ymaps3';
import markImage from '@/assets/images/icons/mark.svg';
import { ref } from 'vue';


const marks = ref([
    {   coordinates: [37.94, 55.76],
        workload: 0 
    },
    {   coordinates: [37.74, 55.76],
        workload: 1
    },
    {   coordinates: [38.14, 55.86],
        workload: 5
    },
    {   coordinates: [37.74, 55.76],
        workload: 4
    },
    {   coordinates: [37.88, 55.70],
        workload: 3
    },
    {   coordinates: [38.02, 55.74],
        workload: 1 
    },
    {   coordinates: [37.68, 55.82],
        workload: 1
    },
    {   coordinates: [37.86, 55.82],
        workload: 5
    },
    {   coordinates: [37.94, 55.76],
        workload: 4
    },
    {   coordinates: [37.82, 55.78],
        workload: 3 
    },
    {   coordinates: [37.90, 55.68],
        workload: 2 
    },
])
const userCoordinates = ref(null)

async function getUserCoordinates() {
    return new Promise((resolve, reject) => {
        navigator.geolocation.getCurrentPosition(
            (position) => {
                const latitude = position.coords.latitude;
                const longitude = position.coords.longitude;
                userCoordinates.value = [longitude, latitude];
                console.log('Получены координаты пользователя:', userCoordinates.value);
                resolve();
            },
            (error) => {
                console.error('Ошибка получения координат пользователя:', error);
                userCoordinates.value = [37.94, 55.76]
                reject(error);
            }
        );
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

        map = new YMap(document.getElementById('map'), {location: LOCATION});
        map.addChild(new YMapDefaultSchemeLayer())
        .addChild(new YMapDefaultFeaturesLayer())
        .addChild(new YMapFeatureDataSource({id: 'my-source'}))
        .addChild(new YMapLayer({source: 'my-source', type: 'markers', zIndex: 1800}))



        const contentPin = document.createElement('div');
        contentPin.classList.add('mark')
        contentPin.innerHTML = `<img src=${markImage}>`;



        createCluster(marks.value)

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
    } catch (err) {
        console.log(err)
    }
}

const createCluster = (marks) => {
    const contentPin = document.createElement('div');
    contentPin.classList.add('mark');



    contentPin.innerHTML = `<img src=${markImage}>`;


    const marker = (feature) =>
        new ymaps3.YMapMarker({
            coordinates: feature.geometry.coordinates,
            source: 'my-source'
        },
        contentPin.cloneNode(true)
    );

    const cluster = (coordinates, features) =>
        new ymaps3.YMapMarker({
            coordinates,
            source: 'my-source'
        },
        circle(features.length).cloneNode(true)
    );

    const points = marks.map((mark, i) => ({
        type: 'Feature',
        id: i,
        geometry: { coordinates: mark.coordinates },
        properties: { name: 'Point of issue of orders', workload: mark.workload }
    }));

    const clusterer = new YMapClusterer({
        method: clusterByGrid({ gridSize: 64 }),
        features: points,
        marker,
        cluster
    });

    map.addChild(clusterer);
    map.update();
}


</script>

<template>
    <div id="map"></div>
</template>


<style lang="scss">
#map {
    width: 100%;
    height: calc(100vh - 71px);
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
</style>