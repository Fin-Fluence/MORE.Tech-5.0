<script setup>
import * as ymaps3 from 'ymaps3';
import markImage from '@/assets/images/icons/mark.svg';
import { ref } from 'vue';

const coordinates = ref([
    [37.64, 55.76],
]);

const coordinates2 = ref([
    [37.94, 55.76],
    [37.74, 55.76],
]);
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
let YMapControls = null;
let YMapClusterer = null;
let clusterByGrid = null;
let YMapGeolocationControl = null;
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
        YMapControls = ymaps3.YMapControls;

        YMapClusterer = await ymaps3.import('@yandex/ymaps3-clusterer@0.0.1');
        clusterByGrid = YMapClusterer.clusterByGrid;
        YMapClusterer = YMapClusterer.YMapClusterer;

        const controlsPackage = await ymaps3.import('@yandex/ymaps3-controls@0.0.1');
        YMapGeolocationControl = controlsPackage.YMapGeolocationControl;

        // Координаты центра карты
        const CENTER_COORDINATES = userCoordinates.value;
        const LOCATION = {center: CENTER_COORDINATES, zoom: 9};

        const controls = new YMapControls({position: 'right'});
        const geolocationControl = new YMapGeolocationControl();
        controls.addChild(geolocationControl);



        map = new YMap(document.getElementById('map'), {location: LOCATION});
        map.addChild(new YMapDefaultSchemeLayer())
        .addChild(new YMapDefaultFeaturesLayer())
        .addChild(new YMapFeatureDataSource({id: 'my-source'}))
        .addChild(new YMapLayer({source: 'my-source', type: 'markers', zIndex: 1800}))
        .addChild(controls);



        const contentPin = document.createElement('div');
        contentPin.classList.add('mark')
        contentPin.innerHTML = `<img src=${markImage}>`;




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

const createCluster = (coordinates) => {
    console.log(coordinates)
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

    const points = coordinates.map((lnglat, i) => ({
        type: 'Feature',
        id: i,
        geometry: { coordinates: lnglat },
        properties: { name: 'Point of issue of orders' }
    }));

    const clusterer = new YMapClusterer({
        method: clusterByGrid({ gridSize: 64 }),
        features: points,
        marker,
        cluster
    });

    map.addChild(clusterer);
    map.update()
}

setTimeout(() => {
    createCluster(coordinates.value)
}, 1000);

setTimeout(() => {
    createCluster(coordinates2.value)
}, 2000);



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