package main

import (
	"bytes"
	"github.com/OlegSchwann/core-services-train/2_categories/category_item"
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"testing"
)

// language=HTML
var inputHtml = []byte(`
<!DOCTYPE html>
<testHtml lang="ru">
<head>
    <meta charset="utf-8">
    <title>Купить авто в России — 695191 объявление </title>
    <meta name="description"
          content="Актуальные предложения на Авито — подержанные, новые авто и запчасти. Фотографии, цены и характеристики, поиск по марке и модели. Проверка истории по VIN.">
</head>
<body>
<div class="js-rubricator">
    <div data-marker="rubricator" class="rubricator-root-nQPOX rubricator-root_border-2t6_-">
        <ul data-marker="rubricator/list" class="rubricator-list-2_NJ4 ">
            <li data-item="data-item" data-visible="data-visible" data-marker="category[7]"
                class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_opened-162dw">
                <div><i data-marker="category[7]/expander" data-category-id="7"
                        class="rubricator-expander-icon-3wgf-"></i><a title="Бытовая электроника"
                                                                      href="/rossiya/bytovaya_elektronika?s_trg=11&amp;cd=1"
                                                                      data-marker="category[7]/current"
                                                                      data-category-id="7"
                                                                      class="rubricator-link-3w6_y   rubricator-link_current-zll0E">Бытовая
                    электроника</a></div>
                <ul data-marker="category[7]/subs" class="rubricator-submenu-18HMk">
                    <li data-item="data-item" data-marker="category[43]"
                        class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                        <div><i data-marker="category[43]/expander" data-category-id="43"
                                class="rubricator-expander-icon-3wgf-"></i><a title="Аудио и видео"
                                                                              href="/rossiya/audio_i_video?s_trg=11&amp;cd=1"
                                                                              data-marker="category[43]/link"
                                                                              data-category-id="43"
                                                                              class="rubricator-link-3w6_y">Аудио и
                            видео</a></div>
                        <ul data-marker="category[43]/subs" class="rubricator-submenu-18HMk">
                            <li data-marker="category[285]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="MP3-плееры" href="/rossiya/audio_i_video/mp3-pleery?s_trg=11&amp;cd=1"
                                        data-marker="category[285]/link" data-category-id="285"
                                        class="rubricator-link-3w6_y">MP3-плееры</a></div>
                            </li>
                            <li data-marker="category[289]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Акустика, колонки, сабвуферы"
                                        href="/rossiya/audio_i_video/akustika_kolonki_sabvufery?s_trg=11&amp;cd=1"
                                        data-marker="category[289]/link" data-category-id="289"
                                        class="rubricator-link-3w6_y">Акустика, колонки, сабвуферы</a></div>
                            </li>
                            <li data-marker="category[284]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Видео, DVD и Blu-ray плееры"
                                        href="/rossiya/audio_i_video/video_dvd_i_blu-ray_pleery?s_trg=11&amp;cd=1"
                                        data-marker="category[284]/link" data-category-id="284"
                                        class="rubricator-link-3w6_y">Видео, DVD и Blu-ray плееры</a></div>
                            </li>
                            <li data-marker="category[288]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Видеокамеры" href="/rossiya/audio_i_video/videokamery?s_trg=11&amp;cd=1"
                                        data-marker="category[288]/link" data-category-id="288"
                                        class="rubricator-link-3w6_y">Видеокамеры</a></div>
                            </li>
                            <li data-marker="category[290]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Кабели и адаптеры"
                                        href="/rossiya/audio_i_video/kabeli_i_adaptery?s_trg=11&amp;cd=1"
                                        data-marker="category[290]/link" data-category-id="290"
                                        class="rubricator-link-3w6_y">Кабели и адаптеры</a></div>
                            </li>
                            <li data-marker="category[291]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Микрофоны" href="/rossiya/audio_i_video/mikrofony?s_trg=11&amp;cd=1"
                                        data-marker="category[291]/link" data-category-id="291"
                                        class="rubricator-link-3w6_y">Микрофоны</a></div>
                            </li>
                            <li data-marker="category[283]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Музыка и фильмы"
                                        href="/rossiya/audio_i_video/muzyka_i_filmy?s_trg=11&amp;cd=1"
                                        data-marker="category[283]/link" data-category-id="283"
                                        class="rubricator-link-3w6_y">Музыка и фильмы</a></div>
                            </li>
                            <li data-marker="category[292]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Музыкальные центры, магнитолы"
                                        href="/rossiya/audio_i_video/muzykalnye_tsentry_magnitoly?s_trg=11&amp;cd=1"
                                        data-marker="category[292]/link" data-category-id="292"
                                        class="rubricator-link-3w6_y">Музыкальные центры, магнитолы</a></div>
                            </li>
                            <li data-marker="category[293]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Наушники" href="/rossiya/audio_i_video/naushniki?s_trg=11&amp;cd=1"
                                        data-marker="category[293]/link" data-category-id="293"
                                        class="rubricator-link-3w6_y">Наушники</a></div>
                            </li>
                            <li data-marker="category[287]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Телевизоры и проекторы"
                                        href="/rossiya/audio_i_video/televizory_i_proektory?s_trg=11&amp;cd=1"
                                        data-marker="category[287]/link" data-category-id="287"
                                        class="rubricator-link-3w6_y">Телевизоры и проекторы</a></div>
                            </li>
                            <li data-marker="category[286]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Усилители и ресиверы"
                                        href="/rossiya/audio_i_video/usiliteli_i_resivery?s_trg=11&amp;cd=1"
                                        data-marker="category[286]/link" data-category-id="286"
                                        class="rubricator-link-3w6_y">Усилители и ресиверы</a></div>
                            </li>
                            <li data-marker="category[282]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Аксессуары" href="/rossiya/audio_i_video/aksessuary?s_trg=11&amp;cd=1"
                                        data-marker="category[282]/link" data-category-id="282"
                                        class="rubricator-link-3w6_y">Аксессуары</a></div>
                            </li>
                        </ul>
                    </li>
                    <li data-item="data-item" data-marker="category[44]"
                        class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                        <div><i data-marker="category[44]/expander" data-category-id="44"
                                class="rubricator-expander-icon-3wgf-"></i><a title="Игры, приставки и программы"
                                                                              href="/rossiya/igry_pristavki_i_programmy?s_trg=11&amp;cd=1"
                                                                              data-marker="category[44]/link"
                                                                              data-category-id="44"
                                                                              class="rubricator-link-3w6_y">Игры,
                            приставки и программы</a></div>
                        <ul data-marker="category[44]/subs" class="rubricator-submenu-18HMk">
                            <li data-marker="category[573]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Игры для приставок"
                                        href="/rossiya/igry_pristavki_i_programmy/igry_dlya_pristavok?s_trg=11&amp;cd=1"
                                        data-marker="category[573]/link" data-category-id="573"
                                        class="rubricator-link-3w6_y">Игры для приставок</a></div>
                            </li>
                            <li data-marker="category[575]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Игровые приставки"
                                        href="/rossiya/igry_pristavki_i_programmy/igrovye_pristavki?s_trg=11&amp;cd=1"
                                        data-marker="category[575]/link" data-category-id="575"
                                        class="rubricator-link-3w6_y">Игровые приставки</a></div>
                            </li>
                            <li data-marker="category[574]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Компьютерные игры"
                                        href="/rossiya/igry_pristavki_i_programmy/kompyuternye_igry?s_trg=11&amp;cd=1"
                                        data-marker="category[574]/link" data-category-id="574"
                                        class="rubricator-link-3w6_y">Компьютерные игры</a></div>
                            </li>
                            <li data-marker="category[572]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Программы"
                                        href="/rossiya/igry_pristavki_i_programmy/programmy?s_trg=11&amp;cd=1"
                                        data-marker="category[572]/link" data-category-id="572"
                                        class="rubricator-link-3w6_y">Программы</a></div>
                            </li>
                        </ul>
                    </li>
                    <li data-marker="category[45]" class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                        <div><a title="Настольные компьютеры" href="/rossiya/nastolnye_kompyutery?s_trg=11&amp;cd=1"
                                data-marker="category[45]/link" data-category-id="45" class="rubricator-link-3w6_y">Настольные
                            компьютеры</a></div>
                    </li>
                    <li data-marker="category[46]" class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                        <div><a title="Ноутбуки" href="/rossiya/noutbuki?s_trg=11&amp;cd=1"
                                data-marker="category[46]/link" data-category-id="46" class="rubricator-link-3w6_y">Ноутбуки</a>
                        </div>
                    </li>
                    <li data-item="data-item" data-marker="category[47]"
                        class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                        <div><i data-marker="category[47]/expander" data-category-id="47"
                                class="rubricator-expander-icon-3wgf-"></i><a title="Оргтехника и расходники"
                                                                              href="/rossiya/orgtehnika_i_rashodniki?s_trg=11&amp;cd=1"
                                                                              data-marker="category[47]/link"
                                                                              data-category-id="47"
                                                                              class="rubricator-link-3w6_y">Оргтехника и
                            расходники</a></div>
                        <ul data-marker="category[47]/subs" class="rubricator-submenu-18HMk">
                            <li data-marker="category[579]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="МФУ, копиры и сканеры"
                                        href="/rossiya/orgtehnika_i_rashodniki/mfu_kopiry_i_skanery?s_trg=11&amp;cd=1"
                                        data-marker="category[579]/link" data-category-id="579"
                                        class="rubricator-link-3w6_y">МФУ, копиры и сканеры</a></div>
                            </li>
                            <li data-marker="category[576]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Принтеры"
                                        href="/rossiya/orgtehnika_i_rashodniki/printery?s_trg=11&amp;cd=1"
                                        data-marker="category[576]/link" data-category-id="576"
                                        class="rubricator-link-3w6_y">Принтеры</a></div>
                            </li>
                            <li data-marker="category[580]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Телефония"
                                        href="/rossiya/orgtehnika_i_rashodniki/telefoniya?s_trg=11&amp;cd=1"
                                        data-marker="category[580]/link" data-category-id="580"
                                        class="rubricator-link-3w6_y">Телефония</a></div>
                            </li>
                            <li data-marker="category[581]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="ИБП, сетевые фильтры"
                                        href="/rossiya/orgtehnika_i_rashodniki/ibp_setevye_filtry?s_trg=11&amp;cd=1"
                                        data-marker="category[581]/link" data-category-id="581"
                                        class="rubricator-link-3w6_y">ИБП, сетевые фильтры</a></div>
                            </li>
                            <li data-marker="category[582]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Уничтожители бумаг"
                                        href="/rossiya/orgtehnika_i_rashodniki/unichtozhiteli_bumag?s_trg=11&amp;cd=1"
                                        data-marker="category[582]/link" data-category-id="582"
                                        class="rubricator-link-3w6_y">Уничтожители бумаг</a></div>
                            </li>
                            <li data-item="data-item" data-marker="category[577]"
                                class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                                <div><i data-marker="category[577]/expander" data-category-id="577"
                                        class="rubricator-expander-icon-3wgf-"></i><a title="Расходные материалы"
                                                                                      href="/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy?s_trg=11&amp;cd=1"
                                                                                      data-marker="category[577]/link"
                                                                                      data-category-id="577"
                                                                                      class="rubricator-link-3w6_y">Расходные
                                    материалы</a></div>
                                <ul data-marker="category[577]/subs" class="rubricator-submenu-18HMk">
                                    <li data-marker="category[3744]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Блоки питания и батареи"
                                                href="/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy/bloki_pitaniya_i_batarei?s_trg=11&amp;cd=1"
                                                data-marker="category[3744]/link" data-category-id="3744"
                                                class="rubricator-link-3w6_y">Блоки питания и батареи</a></div>
                                    </li>
                                    <li data-marker="category[3745]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Болванки"
                                                href="/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy/bolvanki?s_trg=11&amp;cd=1"
                                                data-marker="category[3745]/link" data-category-id="3745"
                                                class="rubricator-link-3w6_y">Болванки</a></div>
                                    </li>
                                    <li data-marker="category[3746]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Бумага"
                                                href="/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy/bumaga?s_trg=11&amp;cd=1"
                                                data-marker="category[3746]/link" data-category-id="3746"
                                                class="rubricator-link-3w6_y">Бумага</a></div>
                                    </li>
                                    <li data-marker="category[3747]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Кабели и адаптеры"
                                                href="/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy/kabeli_i_adaptery?s_trg=11&amp;cd=1"
                                                data-marker="category[3747]/link" data-category-id="3747"
                                                class="rubricator-link-3w6_y">Кабели и адаптеры</a></div>
                                    </li>
                                    <li data-marker="category[3748]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Картриджи"
                                                href="/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy/kartridzhi?s_trg=11&amp;cd=1"
                                                data-marker="category[3748]/link" data-category-id="3748"
                                                class="rubricator-link-3w6_y">Картриджи</a></div>
                                    </li>
                                </ul>
                            </li>
                            <li data-marker="category[578]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Канцелярия"
                                        href="/rossiya/orgtehnika_i_rashodniki/kantselyariya?s_trg=11&amp;cd=1"
                                        data-marker="category[578]/link" data-category-id="578"
                                        class="rubricator-link-3w6_y">Канцелярия</a></div>
                            </li>
                        </ul>
                    </li>
                    <li data-item="data-item" data-marker="category[48]"
                        class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                        <div><i data-marker="category[48]/expander" data-category-id="48"
                                class="rubricator-expander-icon-3wgf-"></i><a title="Планшеты и электронные книги"
                                                                              href="/rossiya/planshety_i_elektronnye_knigi?s_trg=11&amp;cd=1"
                                                                              data-marker="category[48]/link"
                                                                              data-category-id="48"
                                                                              class="rubricator-link-3w6_y">Планшеты и
                            электронные книги</a></div>
                        <ul data-marker="category[48]/subs" class="rubricator-submenu-18HMk">
                            <li data-marker="category[569]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Планшеты"
                                        href="/rossiya/planshety_i_elektronnye_knigi/planshety?s_trg=11&amp;cd=1"
                                        data-marker="category[569]/link" data-category-id="569"
                                        class="rubricator-link-3w6_y">Планшеты</a></div>
                            </li>
                            <li data-marker="category[570]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Электронные книги"
                                        href="/rossiya/planshety_i_elektronnye_knigi/elektronnye_knigi?s_trg=11&amp;cd=1"
                                        data-marker="category[570]/link" data-category-id="570"
                                        class="rubricator-link-3w6_y">Электронные книги</a></div>
                            </li>
                            <li data-item="data-item" data-marker="category[571]"
                                class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                                <div><i data-marker="category[571]/expander" data-category-id="571"
                                        class="rubricator-expander-icon-3wgf-"></i><a title="Аксессуары"
                                                                                      href="/rossiya/planshety_i_elektronnye_knigi/aksessuary?s_trg=11&amp;cd=1"
                                                                                      data-marker="category[571]/link"
                                                                                      data-category-id="571"
                                                                                      class="rubricator-link-3w6_y">Аксессуары</a>
                                </div>
                                <ul data-marker="category[571]/subs" class="rubricator-submenu-18HMk">
                                    <li data-marker="category[3736]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Аккумуляторы"
                                                href="/rossiya/planshety_i_elektronnye_knigi/aksessuary/akkumulyatory?s_trg=11&amp;cd=1"
                                                data-marker="category[3736]/link" data-category-id="3736"
                                                class="rubricator-link-3w6_y">Аккумуляторы</a></div>
                                    </li>
                                    <li data-marker="category[3737]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Гарнитуры и наушники"
                                                href="/rossiya/planshety_i_elektronnye_knigi/aksessuary/garnitury_i_naushniki?s_trg=11&amp;cd=1"
                                                data-marker="category[3737]/link" data-category-id="3737"
                                                class="rubricator-link-3w6_y">Гарнитуры и наушники</a></div>
                                    </li>
                                    <li data-marker="category[3738]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Док-станции"
                                                href="/rossiya/planshety_i_elektronnye_knigi/aksessuary/dok-stantsii?s_trg=11&amp;cd=1"
                                                data-marker="category[3738]/link" data-category-id="3738"
                                                class="rubricator-link-3w6_y">Док-станции</a></div>
                                    </li>
                                    <li data-marker="category[3739]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Зарядные устройства"
                                                href="/rossiya/planshety_i_elektronnye_knigi/aksessuary/zaryadnye_ustroystva?s_trg=11&amp;cd=1"
                                                data-marker="category[3739]/link" data-category-id="3739"
                                                class="rubricator-link-3w6_y">Зарядные устройства</a></div>
                                    </li>
                                    <li data-marker="category[3740]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Кабели и адаптеры"
                                                href="/rossiya/planshety_i_elektronnye_knigi/aksessuary/kabeli_i_adaptery?s_trg=11&amp;cd=1"
                                                data-marker="category[3740]/link" data-category-id="3740"
                                                class="rubricator-link-3w6_y">Кабели и адаптеры</a></div>
                                    </li>
                                    <li data-marker="category[3741]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Модемы и роутеры"
                                                href="/rossiya/planshety_i_elektronnye_knigi/aksessuary/modemy_i_routery?s_trg=11&amp;cd=1"
                                                data-marker="category[3741]/link" data-category-id="3741"
                                                class="rubricator-link-3w6_y">Модемы и роутеры</a></div>
                                    </li>
                                    <li data-marker="category[3742]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Стилусы"
                                                href="/rossiya/planshety_i_elektronnye_knigi/aksessuary/stilusy?s_trg=11&amp;cd=1"
                                                data-marker="category[3742]/link" data-category-id="3742"
                                                class="rubricator-link-3w6_y">Стилусы</a></div>
                                    </li>
                                    <li data-marker="category[3743]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Чехлы и плёнки"
                                                href="/rossiya/planshety_i_elektronnye_knigi/aksessuary/chehly_i_plenki?s_trg=11&amp;cd=1"
                                                data-marker="category[3743]/link" data-category-id="3743"
                                                class="rubricator-link-3w6_y">Чехлы и плёнки</a></div>
                                    </li>
                                    <li data-marker="category[3735]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Другое"
                                                href="/rossiya/planshety_i_elektronnye_knigi/aksessuary/drugoe?s_trg=11&amp;cd=1"
                                                data-marker="category[3735]/link" data-category-id="3735"
                                                class="rubricator-link-3w6_y">Другое</a></div>
                                    </li>
                                </ul>
                            </li>
                        </ul>
                    </li>
                    <li data-item="data-item" data-marker="category[49]"
                        class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                        <div><i data-marker="category[49]/expander" data-category-id="49"
                                class="rubricator-expander-icon-3wgf-"></i><a title="Телефоны"
                                                                              href="/rossiya/telefony?s_trg=11&amp;cd=1"
                                                                              data-marker="category[49]/link"
                                                                              data-category-id="49"
                                                                              class="rubricator-link-3w6_y">Телефоны</a>
                        </div>
                        <ul data-marker="category[49]/subs" class="rubricator-submenu-18HMk">
                            <li data-marker="category[410]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Acer" href="/rossiya/telefony/acer?s_trg=11&amp;cd=1"
                                        data-marker="category[410]/link" data-category-id="410"
                                        class="rubricator-link-3w6_y">Acer</a></div>
                            </li>
                            <li data-marker="category[411]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Alcatel" href="/rossiya/telefony/alcatel?s_trg=11&amp;cd=1"
                                        data-marker="category[411]/link" data-category-id="411"
                                        class="rubricator-link-3w6_y">Alcatel</a></div>
                            </li>
                            <li data-marker="category[84907]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="ASUS" href="/rossiya/telefony/asus?s_trg=11&amp;cd=1"
                                        data-marker="category[84907]/link" data-category-id="84907"
                                        class="rubricator-link-3w6_y">ASUS</a></div>
                            </li>
                            <li data-marker="category[405]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="BlackBerry" href="/rossiya/telefony/blackberry?s_trg=11&amp;cd=1"
                                        data-marker="category[405]/link" data-category-id="405"
                                        class="rubricator-link-3w6_y">BlackBerry</a></div>
                            </li>
                            <li data-marker="category[84917]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="BQ" href="/rossiya/telefony/bq?s_trg=11&amp;cd=1"
                                        data-marker="category[84917]/link" data-category-id="84917"
                                        class="rubricator-link-3w6_y">BQ</a></div>
                            </li>
                            <li data-marker="category[84908]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="DEXP" href="/rossiya/telefony/dexp?s_trg=11&amp;cd=1"
                                        data-marker="category[84908]/link" data-category-id="84908"
                                        class="rubricator-link-3w6_y">DEXP</a></div>
                            </li>
                            <li data-marker="category[414]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Explay" href="/rossiya/telefony/explay?s_trg=11&amp;cd=1"
                                        data-marker="category[414]/link" data-category-id="414"
                                        class="rubricator-link-3w6_y">Explay</a></div>
                            </li>
                            <li data-marker="category[408]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Fly" href="/rossiya/telefony/fly?s_trg=11&amp;cd=1"
                                        data-marker="category[408]/link" data-category-id="408"
                                        class="rubricator-link-3w6_y">Fly</a></div>
                            </li>
                            <li data-marker="category[84909]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Highscreen" href="/rossiya/telefony/highscreen?s_trg=11&amp;cd=1"
                                        data-marker="category[84909]/link" data-category-id="84909"
                                        class="rubricator-link-3w6_y">Highscreen</a></div>
                            </li>
                            <li data-marker="category[406]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="HTC" href="/rossiya/telefony/htc?s_trg=11&amp;cd=1"
                                        data-marker="category[406]/link" data-category-id="406"
                                        class="rubricator-link-3w6_y">HTC</a></div>
                            </li>
                            <li data-marker="category[415]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Huawei" href="/rossiya/telefony/huawei?s_trg=11&amp;cd=1"
                                        data-marker="category[415]/link" data-category-id="415"
                                        class="rubricator-link-3w6_y">Huawei</a></div>
                            </li>
                            <li data-marker="category[398]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="iPhone" href="/rossiya/telefony/iphone?s_trg=11&amp;cd=1"
                                        data-marker="category[398]/link" data-category-id="398"
                                        class="rubricator-link-3w6_y">iPhone</a></div>
                            </li>
                            <li data-marker="category[416]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Lenovo" href="/rossiya/telefony/lenovo?s_trg=11&amp;cd=1"
                                        data-marker="category[416]/link" data-category-id="416"
                                        class="rubricator-link-3w6_y">Lenovo</a></div>
                            </li>
                            <li data-marker="category[399]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="LG" href="/rossiya/telefony/lg?s_trg=11&amp;cd=1"
                                        data-marker="category[399]/link" data-category-id="399"
                                        class="rubricator-link-3w6_y">LG</a></div>
                            </li>
                            <li data-marker="category[84910]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Meizu" href="/rossiya/telefony/meizu?s_trg=11&amp;cd=1"
                                        data-marker="category[84910]/link" data-category-id="84910"
                                        class="rubricator-link-3w6_y">Meizu</a></div>
                            </li>
                            <li data-marker="category[84911]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Micromax" href="/rossiya/telefony/micromax?s_trg=11&amp;cd=1"
                                        data-marker="category[84911]/link" data-category-id="84911"
                                        class="rubricator-link-3w6_y">Micromax</a></div>
                            </li>
                            <li data-marker="category[84916]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Microsoft" href="/rossiya/telefony/microsoft?s_trg=11&amp;cd=1"
                                        data-marker="category[84916]/link" data-category-id="84916"
                                        class="rubricator-link-3w6_y">Microsoft</a></div>
                            </li>
                            <li data-marker="category[400]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Motorola" href="/rossiya/telefony/motorola?s_trg=11&amp;cd=1"
                                        data-marker="category[400]/link" data-category-id="400"
                                        class="rubricator-link-3w6_y">Motorola</a></div>
                            </li>
                            <li data-marker="category[417]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="MTS" href="/rossiya/telefony/mts?s_trg=11&amp;cd=1"
                                        data-marker="category[417]/link" data-category-id="417"
                                        class="rubricator-link-3w6_y">MTS</a></div>
                            </li>
                            <li data-marker="category[401]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Nokia" href="/rossiya/telefony/nokia?s_trg=11&amp;cd=1"
                                        data-marker="category[401]/link" data-category-id="401"
                                        class="rubricator-link-3w6_y">Nokia</a></div>
                            </li>
                            <li data-marker="category[84912]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Panasonic" href="/rossiya/telefony/panasonic?s_trg=11&amp;cd=1"
                                        data-marker="category[84912]/link" data-category-id="84912"
                                        class="rubricator-link-3w6_y">Panasonic</a></div>
                            </li>
                            <li data-marker="category[409]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Philips" href="/rossiya/telefony/philips?s_trg=11&amp;cd=1"
                                        data-marker="category[409]/link" data-category-id="409"
                                        class="rubricator-link-3w6_y">Philips</a></div>
                            </li>
                            <li data-marker="category[84913]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Prestigio" href="/rossiya/telefony/prestigio?s_trg=11&amp;cd=1"
                                        data-marker="category[84913]/link" data-category-id="84913"
                                        class="rubricator-link-3w6_y">Prestigio</a></div>
                            </li>
                            <li data-marker="category[402]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Samsung" href="/rossiya/telefony/samsung?s_trg=11&amp;cd=1"
                                        data-marker="category[402]/link" data-category-id="402"
                                        class="rubricator-link-3w6_y">Samsung</a></div>
                            </li>
                            <li data-marker="category[403]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Siemens" href="/rossiya/telefony/siemens?s_trg=11&amp;cd=1"
                                        data-marker="category[403]/link" data-category-id="403"
                                        class="rubricator-link-3w6_y">Siemens</a></div>
                            </li>
                            <li data-marker="category[412]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="SkyLink" href="/rossiya/telefony/skylink?s_trg=11&amp;cd=1"
                                        data-marker="category[412]/link" data-category-id="412"
                                        class="rubricator-link-3w6_y">SkyLink</a></div>
                            </li>
                            <li data-marker="category[404]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Sony" href="/rossiya/telefony/sony?s_trg=11&amp;cd=1"
                                        data-marker="category[404]/link" data-category-id="404"
                                        class="rubricator-link-3w6_y">Sony</a></div>
                            </li>
                            <li data-marker="category[84914]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="teXet" href="/rossiya/telefony/texet?s_trg=11&amp;cd=1"
                                        data-marker="category[84914]/link" data-category-id="84914"
                                        class="rubricator-link-3w6_y">teXet</a></div>
                            </li>
                            <li data-marker="category[407]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Vertu" href="/rossiya/telefony/vertu?s_trg=11&amp;cd=1"
                                        data-marker="category[407]/link" data-category-id="407"
                                        class="rubricator-link-3w6_y">Vertu</a></div>
                            </li>
                            <li data-marker="category[84915]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Xiaomi" href="/rossiya/telefony/xiaomi?s_trg=11&amp;cd=1"
                                        data-marker="category[84915]/link" data-category-id="84915"
                                        class="rubricator-link-3w6_y">Xiaomi</a></div>
                            </li>
                            <li data-marker="category[418]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="ZTE" href="/rossiya/telefony/zte?s_trg=11&amp;cd=1"
                                        data-marker="category[418]/link" data-category-id="418"
                                        class="rubricator-link-3w6_y">ZTE</a></div>
                            </li>
                            <li data-marker="category[397]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Другие марки" href="/rossiya/telefony/drugie_marki?s_trg=11&amp;cd=1"
                                        data-marker="category[397]/link" data-category-id="397"
                                        class="rubricator-link-3w6_y">Другие марки</a></div>
                            </li>
                            <li data-marker="category[394]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Рации" href="/rossiya/telefony/ratsii?s_trg=11&amp;cd=1"
                                        data-marker="category[394]/link" data-category-id="394"
                                        class="rubricator-link-3w6_y">Рации</a></div>
                            </li>
                            <li data-marker="category[396]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND  rubricator-item_closed-12NK-">
                                <div><a title="Стационарные телефоны"
                                        href="/rossiya/telefony/statsionarnye_telefony?s_trg=11&amp;cd=1"
                                        data-marker="category[396]/link" data-category-id="396"
                                        class="rubricator-link-3w6_y">Стационарные телефоны</a></div>
                            </li>
                            <li data-item="data-item" data-marker="category[413]"
                                class="rubricator-item-31NYs js-more-category rubricator-more-category-y3f_q rubricator-more-category_hidden-3b7ND rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                                <div><i data-marker="category[413]/expander" data-category-id="413"
                                        class="rubricator-expander-icon-3wgf-"></i><a title="Аксессуары"
                                                                                      href="/rossiya/telefony/aksessuary?s_trg=11&amp;cd=1"
                                                                                      data-marker="category[413]/link"
                                                                                      data-category-id="413"
                                                                                      class="rubricator-link-3w6_y">Аксессуары</a>
                                </div>
                                <ul data-marker="category[413]/subs" class="rubricator-submenu-18HMk">
                                    <li data-marker="category[3729]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Аккумуляторы"
                                                href="/rossiya/telefony/aksessuary/akkumulyatory?s_trg=11&amp;cd=1"
                                                data-marker="category[3729]/link" data-category-id="3729"
                                                class="rubricator-link-3w6_y">Аккумуляторы</a></div>
                                    </li>
                                    <li data-marker="category[3730]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Гарнитуры и наушники"
                                                href="/rossiya/telefony/aksessuary/garnitury_i_naushniki?s_trg=11&amp;cd=1"
                                                data-marker="category[3730]/link" data-category-id="3730"
                                                class="rubricator-link-3w6_y">Гарнитуры и наушники</a></div>
                                    </li>
                                    <li data-marker="category[3731]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Зарядные устройства"
                                                href="/rossiya/telefony/aksessuary/zaryadnye_ustroystva?s_trg=11&amp;cd=1"
                                                data-marker="category[3731]/link" data-category-id="3731"
                                                class="rubricator-link-3w6_y">Зарядные устройства</a></div>
                                    </li>
                                    <li data-marker="category[3732]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Кабели и адаптеры"
                                                href="/rossiya/telefony/aksessuary/kabeli_i_adaptery?s_trg=11&amp;cd=1"
                                                data-marker="category[3732]/link" data-category-id="3732"
                                                class="rubricator-link-3w6_y">Кабели и адаптеры</a></div>
                                    </li>
                                    <li data-marker="category[3733]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Модемы и роутеры"
                                                href="/rossiya/telefony/aksessuary/modemy_i_routery?s_trg=11&amp;cd=1"
                                                data-marker="category[3733]/link" data-category-id="3733"
                                                class="rubricator-link-3w6_y">Модемы и роутеры</a></div>
                                    </li>
                                    <li data-marker="category[3734]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Чехлы и плёнки"
                                                href="/rossiya/telefony/aksessuary/chehly_i_plenki?s_trg=11&amp;cd=1"
                                                data-marker="category[3734]/link" data-category-id="3734"
                                                class="rubricator-link-3w6_y">Чехлы и плёнки</a></div>
                                    </li>
                                    <li data-marker="category[3728]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Запчасти"
                                                href="/rossiya/telefony/aksessuary/zapchasti?s_trg=11&amp;cd=1"
                                                data-marker="category[3728]/link" data-category-id="3728"
                                                class="rubricator-link-3w6_y">Запчасти</a></div>
                                    </li>
                                </ul>
                            </li>

                            <button data-marker="category[49]/more" data-category-id="49"
                                    class="rubricator-button-2w_8d">Ещё
                            </button>
                        </ul>
                    </li>
                    <li data-item="data-item" data-marker="category[42]"
                        class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                        <div><i data-marker="category[42]/expander" data-category-id="42"
                                class="rubricator-expander-icon-3wgf-"></i><a title="Товары для компьютера"
                                                                              href="/rossiya/tovary_dlya_kompyutera?s_trg=11&amp;cd=1"
                                                                              data-marker="category[42]/link"
                                                                              data-category-id="42"
                                                                              class="rubricator-link-3w6_y">Товары для
                            компьютера</a></div>
                        <ul data-marker="category[42]/subs" class="rubricator-submenu-18HMk">
                            <li data-marker="category[593]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Акустика"
                                        href="/rossiya/tovary_dlya_kompyutera/akustika?s_trg=11&amp;cd=1"
                                        data-marker="category[593]/link" data-category-id="593"
                                        class="rubricator-link-3w6_y">Акустика</a></div>
                            </li>
                            <li data-marker="category[588]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Веб-камеры"
                                        href="/rossiya/tovary_dlya_kompyutera/veb-kamery?s_trg=11&amp;cd=1"
                                        data-marker="category[588]/link" data-category-id="588"
                                        class="rubricator-link-3w6_y">Веб-камеры</a></div>
                            </li>
                            <li data-marker="category[589]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Джойстики и рули"
                                        href="/rossiya/tovary_dlya_kompyutera/dzhoystiki_i_ruli?s_trg=11&amp;cd=1"
                                        data-marker="category[589]/link" data-category-id="589"
                                        class="rubricator-link-3w6_y">Джойстики и рули</a></div>
                            </li>
                            <li data-marker="category[587]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Клавиатуры и мыши"
                                        href="/rossiya/tovary_dlya_kompyutera/klaviatury_i_myshi?s_trg=11&amp;cd=1"
                                        data-marker="category[587]/link" data-category-id="587"
                                        class="rubricator-link-3w6_y">Клавиатуры и мыши</a></div>
                            </li>
                            <li data-item="data-item" data-marker="category[591]"
                                class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                                <div><i data-marker="category[591]/expander" data-category-id="591"
                                        class="rubricator-expander-icon-3wgf-"></i><a title="Комплектующие"
                                                                                      href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie?s_trg=11&amp;cd=1"
                                                                                      data-marker="category[591]/link"
                                                                                      data-category-id="591"
                                                                                      class="rubricator-link-3w6_y">Комплектующие</a>
                                </div>
                                <ul data-marker="category[591]/subs" class="rubricator-submenu-18HMk">
                                    <li data-marker="category[3839]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="CD, DVD и Blu-ray приводы"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/cd_dvd_i_blu-ray_privody?s_trg=11&amp;cd=1"
                                                data-marker="category[3839]/link" data-category-id="3839"
                                                class="rubricator-link-3w6_y">CD, DVD и Blu-ray приводы</a></div>
                                    </li>
                                    <li data-marker="category[3840]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Блоки питания"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/bloki_pitaniya?s_trg=11&amp;cd=1"
                                                data-marker="category[3840]/link" data-category-id="3840"
                                                class="rubricator-link-3w6_y">Блоки питания</a></div>
                                    </li>
                                    <li data-marker="category[3841]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Видеокарты"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/videokarty?s_trg=11&amp;cd=1"
                                                data-marker="category[3841]/link" data-category-id="3841"
                                                class="rubricator-link-3w6_y">Видеокарты</a></div>
                                    </li>
                                    <li data-marker="category[3849]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Жёсткие диски"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/zhestkie_diski?s_trg=11&amp;cd=1"
                                                data-marker="category[3849]/link" data-category-id="3849"
                                                class="rubricator-link-3w6_y">Жёсткие диски</a></div>
                                    </li>
                                    <li data-marker="category[3848]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Звуковые карты"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/zvukovye_karty?s_trg=11&amp;cd=1"
                                                data-marker="category[3848]/link" data-category-id="3848"
                                                class="rubricator-link-3w6_y">Звуковые карты</a></div>
                                    </li>
                                    <li data-marker="category[3843]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Контроллеры"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/kontrollery?s_trg=11&amp;cd=1"
                                                data-marker="category[3843]/link" data-category-id="3843"
                                                class="rubricator-link-3w6_y">Контроллеры</a></div>
                                    </li>
                                    <li data-marker="category[3847]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Корпусы"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/korpusy?s_trg=11&amp;cd=1"
                                                data-marker="category[3847]/link" data-category-id="3847"
                                                class="rubricator-link-3w6_y">Корпусы</a></div>
                                    </li>
                                    <li data-marker="category[3846]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Материнские платы"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/materinskie_platy?s_trg=11&amp;cd=1"
                                                data-marker="category[3846]/link" data-category-id="3846"
                                                class="rubricator-link-3w6_y">Материнские платы</a></div>
                                    </li>
                                    <li data-marker="category[3845]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Оперативная память"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/operativnaya_pamyat?s_trg=11&amp;cd=1"
                                                data-marker="category[3845]/link" data-category-id="3845"
                                                class="rubricator-link-3w6_y">Оперативная память</a></div>
                                    </li>
                                    <li data-marker="category[3844]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Процессоры"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/protsessory?s_trg=11&amp;cd=1"
                                                data-marker="category[3844]/link" data-category-id="3844"
                                                class="rubricator-link-3w6_y">Процессоры</a></div>
                                    </li>
                                    <li data-marker="category[3842]"
                                        class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                        <div><a title="Системы охлаждения"
                                                href="/rossiya/tovary_dlya_kompyutera/komplektuyuschie/sistemy_ohlazhdeniya?s_trg=11&amp;cd=1"
                                                data-marker="category[3842]/link" data-category-id="3842"
                                                class="rubricator-link-3w6_y">Системы охлаждения</a></div>
                                    </li>
                                </ul>
                            </li>
                            <li data-marker="category[583]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Мониторы"
                                        href="/rossiya/tovary_dlya_kompyutera/monitory?s_trg=11&amp;cd=1"
                                        data-marker="category[583]/link" data-category-id="583"
                                        class="rubricator-link-3w6_y">Мониторы</a></div>
                            </li>
                            <li data-marker="category[586]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Переносные жёсткие диски"
                                        href="/rossiya/tovary_dlya_kompyutera/perenosnye_zhestkie_diski?s_trg=11&amp;cd=1"
                                        data-marker="category[586]/link" data-category-id="586"
                                        class="rubricator-link-3w6_y">Переносные жёсткие диски</a></div>
                            </li>
                            <li data-marker="category[592]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Сетевое оборудование"
                                        href="/rossiya/tovary_dlya_kompyutera/setevoe_oborudovanie?s_trg=11&amp;cd=1"
                                        data-marker="category[592]/link" data-category-id="592"
                                        class="rubricator-link-3w6_y">Сетевое оборудование</a></div>
                            </li>
                            <li data-marker="category[585]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="ТВ-тюнеры"
                                        href="/rossiya/tovary_dlya_kompyutera/tv-tyunery?s_trg=11&amp;cd=1"
                                        data-marker="category[585]/link" data-category-id="585"
                                        class="rubricator-link-3w6_y">ТВ-тюнеры</a></div>
                            </li>
                            <li data-marker="category[584]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Флэшки и карты памяти"
                                        href="/rossiya/tovary_dlya_kompyutera/fleshki_i_karty_pamyati?s_trg=11&amp;cd=1"
                                        data-marker="category[584]/link" data-category-id="584"
                                        class="rubricator-link-3w6_y">Флэшки и карты памяти</a></div>
                            </li>
                            <li data-marker="category[590]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Аксессуары"
                                        href="/rossiya/tovary_dlya_kompyutera/aksessuary?s_trg=11&amp;cd=1"
                                        data-marker="category[590]/link" data-category-id="590"
                                        class="rubricator-link-3w6_y">Аксессуары</a></div>
                            </li>
                        </ul>
                    </li>
                    <li data-item="data-item" data-marker="category[50]"
                        class="rubricator-item-31NYs  rubricator-item-expandable-2JTLe rubricator-item_closed-12NK-">
                        <div><i data-marker="category[50]/expander" data-category-id="50"
                                class="rubricator-expander-icon-3wgf-"></i><a title="Фототехника"
                                                                              href="/rossiya/fototehnika?s_trg=11&amp;cd=1"
                                                                              data-marker="category[50]/link"
                                                                              data-category-id="50"
                                                                              class="rubricator-link-3w6_y">Фототехника</a>
                        </div>
                        <ul data-marker="category[50]/subs" class="rubricator-submenu-18HMk">
                            <li data-marker="category[603]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Компактные фотоаппараты"
                                        href="/rossiya/fototehnika/kompaktnye_fotoapparaty?s_trg=11&amp;cd=1"
                                        data-marker="category[603]/link" data-category-id="603"
                                        class="rubricator-link-3w6_y">Компактные фотоаппараты</a></div>
                            </li>
                            <li data-marker="category[602]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Зеркальные фотоаппараты"
                                        href="/rossiya/fototehnika/zerkalnye_fotoapparaty?s_trg=11&amp;cd=1"
                                        data-marker="category[602]/link" data-category-id="602"
                                        class="rubricator-link-3w6_y">Зеркальные фотоаппараты</a></div>
                            </li>
                            <li data-marker="category[601]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Плёночные фотоаппараты"
                                        href="/rossiya/fototehnika/plenochnye_fotoapparaty?s_trg=11&amp;cd=1"
                                        data-marker="category[601]/link" data-category-id="601"
                                        class="rubricator-link-3w6_y">Плёночные фотоаппараты</a></div>
                            </li>
                            <li data-marker="category[599]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Бинокли и телескопы"
                                        href="/rossiya/fototehnika/binokli_i_teleskopy?s_trg=11&amp;cd=1"
                                        data-marker="category[599]/link" data-category-id="599"
                                        class="rubricator-link-3w6_y">Бинокли и телескопы</a></div>
                            </li>
                            <li data-marker="category[600]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Объективы" href="/rossiya/fototehnika/obektivy?s_trg=11&amp;cd=1"
                                        data-marker="category[600]/link" data-category-id="600"
                                        class="rubricator-link-3w6_y">Объективы</a></div>
                            </li>
                            <li data-marker="category[598]"
                                class="rubricator-item-31NYs   rubricator-item_closed-12NK-">
                                <div><a title="Оборудование и аксессуары"
                                        href="/rossiya/fototehnika/oborudovanie_i_aksessuary?s_trg=11&amp;cd=1"
                                        data-marker="category[598]/link" data-category-id="598"
                                        class="rubricator-link-3w6_y">Оборудование и аксессуары</a></div>
                            </li>
                        </ul>
                    </li>
                </ul>
            </li>
        </ul>
    </div>
</div>
</body>
</testHtml>
`)

var expectedOutput = category_item.CategoryItems{{
	Title: "Бытовая электроника",
	Href:  "/rossiya/bytovaya_elektronika?s_trg=11&cd=1",
	Childes: []category_item.CategoryItem{{
		Title: "Аудио и видео",
		Href:  "/rossiya/audio_i_video?s_trg=11&cd=1",
		Childes: []category_item.CategoryItem{{
			Title: "MP3-плееры",
			Href:  "/rossiya/audio_i_video/mp3-pleery?s_trg=11&cd=1",
		}, {
			Title: "Акустика, колонки, сабвуферы",
			Href:  "/rossiya/audio_i_video/akustika_kolonki_sabvufery?s_trg=11&cd=1",
		}, {
			Title: "Видео, DVD и Blu-ray плееры",
			Href:  "/rossiya/audio_i_video/video_dvd_i_blu-ray_pleery?s_trg=11&cd=1",
		}, {
			Title: "Видеокамеры",
			Href:  "/rossiya/audio_i_video/videokamery?s_trg=11&cd=1",
		}, {
			Title: "Кабели и адаптеры",
			Href:  "/rossiya/audio_i_video/kabeli_i_adaptery?s_trg=11&cd=1",
		}, {
			Title: "Микрофоны",
			Href:  "/rossiya/audio_i_video/mikrofony?s_trg=11&cd=1",
		}, {
			Title: "Музыка и фильмы",
			Href:  "/rossiya/audio_i_video/muzyka_i_filmy?s_trg=11&cd=1",
		}, {
			Title: "Музыкальные центры, магнитолы",
			Href:  "/rossiya/audio_i_video/muzykalnye_tsentry_magnitoly?s_trg=11&cd=1",
		}, {
			Title: "Наушники",
			Href:  "/rossiya/audio_i_video/naushniki?s_trg=11&cd=1",
		}, {
			Title: "Телевизоры и проекторы",
			Href:  "/rossiya/audio_i_video/televizory_i_proektory?s_trg=11&cd=1",
		}, {
			Title: "Усилители и ресиверы",
			Href:  "/rossiya/audio_i_video/usiliteli_i_resivery?s_trg=11&cd=1",
		}, {
			Title: "Аксессуары",
			Href:  "/rossiya/audio_i_video/aksessuary?s_trg=11&cd=1"}},
	}, {
		Title: "Игры, приставки и программы",
		Href:  "/rossiya/igry_pristavki_i_programmy?s_trg=11&cd=1",
		Childes: []category_item.CategoryItem{{
			Title: "Игры для приставок",
			Href:  "/rossiya/igry_pristavki_i_programmy/igry_dlya_pristavok?s_trg=11&cd=1",
		}, {
			Title: "Игровые приставки",
			Href:  "/rossiya/igry_pristavki_i_programmy/igrovye_pristavki?s_trg=11&cd=1",
		}, {
			Title: "Компьютерные игры",
			Href:  "/rossiya/igry_pristavki_i_programmy/kompyuternye_igry?s_trg=11&cd=1",
		}, {
			Title: "Программы",
			Href:  "/rossiya/igry_pristavki_i_programmy/programmy?s_trg=11&cd=1"}},
	}, {
		Title: "Настольные компьютеры",
		Href:  "/rossiya/nastolnye_kompyutery?s_trg=11&cd=1",
	}, {
		Title: "Ноутбуки",
		Href:  "/rossiya/noutbuki?s_trg=11&cd=1",
	}, {
		Title: "Оргтехника и расходники",
		Href:  "/rossiya/orgtehnika_i_rashodniki?s_trg=11&cd=1",
		Childes: []category_item.CategoryItem{{
			Title: "МФУ, копиры и сканеры",
			Href:  "/rossiya/orgtehnika_i_rashodniki/mfu_kopiry_i_skanery?s_trg=11&cd=1",
		}, {
			Title: "Принтеры",
			Href:  "/rossiya/orgtehnika_i_rashodniki/printery?s_trg=11&cd=1",
		}, {
			Title: "Телефония",
			Href:  "/rossiya/orgtehnika_i_rashodniki/telefoniya?s_trg=11&cd=1",
		}, {
			Title: "ИБП, сетевые фильтры",
			Href:  "/rossiya/orgtehnika_i_rashodniki/ibp_setevye_filtry?s_trg=11&cd=1",
		}, {
			Title: "Уничтожители бумаг",
			Href:  "/rossiya/orgtehnika_i_rashodniki/unichtozhiteli_bumag?s_trg=11&cd=1",
		}, {
			Title: "Расходные материалы",
			Href:  "/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy?s_trg=11&cd=1",
			Childes: []category_item.CategoryItem{{
				Title: "Блоки питания и батареи",
				Href:  "/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy/bloki_pitaniya_i_batarei?s_trg=11&cd=1",
			}, {
				Title: "Болванки",
				Href:  "/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy/bolvanki?s_trg=11&cd=1",
			}, {
				Title: "Бумага",
				Href:  "/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy/bumaga?s_trg=11&cd=1",
			}, {
				Title: "Кабели и адаптеры",
				Href:  "/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy/kabeli_i_adaptery?s_trg=11&cd=1",
			}, {
				Title: "Картриджи",
				Href:  "/rossiya/orgtehnika_i_rashodniki/rashodnye_materialy/kartridzhi?s_trg=11&cd=1"}},
		}, {
			Title: "Канцелярия",
			Href:  "/rossiya/orgtehnika_i_rashodniki/kantselyariya?s_trg=11&cd=1"}},
	}, {
		Title: "Планшеты и электронные книги",
		Href:  "/rossiya/planshety_i_elektronnye_knigi?s_trg=11&cd=1",
		Childes: []category_item.CategoryItem{{
			Title: "Планшеты",
			Href:  "/rossiya/planshety_i_elektronnye_knigi/planshety?s_trg=11&cd=1",
		}, {
			Title: "Электронные книги",
			Href:  "/rossiya/planshety_i_elektronnye_knigi/elektronnye_knigi?s_trg=11&cd=1",
		}, {
			Title: "Аксессуары",
			Href:  "/rossiya/planshety_i_elektronnye_knigi/aksessuary?s_trg=11&cd=1",
			Childes: []category_item.CategoryItem{{
				Title: "Аккумуляторы",
				Href:  "/rossiya/planshety_i_elektronnye_knigi/aksessuary/akkumulyatory?s_trg=11&cd=1",
			}, {
				Title: "Гарнитуры и наушники",
				Href:  "/rossiya/planshety_i_elektronnye_knigi/aksessuary/garnitury_i_naushniki?s_trg=11&cd=1",
			}, {
				Title: "Док-станции",
				Href:  "/rossiya/planshety_i_elektronnye_knigi/aksessuary/dok-stantsii?s_trg=11&cd=1",
			}, {
				Title: "Зарядные устройства",
				Href:  "/rossiya/planshety_i_elektronnye_knigi/aksessuary/zaryadnye_ustroystva?s_trg=11&cd=1",
			}, {
				Title: "Кабели и адаптеры",
				Href:  "/rossiya/planshety_i_elektronnye_knigi/aksessuary/kabeli_i_adaptery?s_trg=11&cd=1",
			}, {
				Title: "Модемы и роутеры",
				Href:  "/rossiya/planshety_i_elektronnye_knigi/aksessuary/modemy_i_routery?s_trg=11&cd=1",
			}, {
				Title: "Стилусы",
				Href:  "/rossiya/planshety_i_elektronnye_knigi/aksessuary/stilusy?s_trg=11&cd=1",
			}, {
				Title: "Чехлы и плёнки",
				Href:  "/rossiya/planshety_i_elektronnye_knigi/aksessuary/chehly_i_plenki?s_trg=11&cd=1",
			}, {
				Title: "Другое",
				Href:  "/rossiya/planshety_i_elektronnye_knigi/aksessuary/drugoe?s_trg=11&cd=1"}}}},
	}, {
		Title: "Телефоны",
		Href:  "/rossiya/telefony?s_trg=11&cd=1",
		Childes: []category_item.CategoryItem{{
			Title: "Acer",
			Href:  "/rossiya/telefony/acer?s_trg=11&cd=1",
		}, {
			Title: "Alcatel",
			Href:  "/rossiya/telefony/alcatel?s_trg=11&cd=1",
		}, {
			Title: "ASUS",
			Href:  "/rossiya/telefony/asus?s_trg=11&cd=1",
		}, {
			Title: "BlackBerry",
			Href:  "/rossiya/telefony/blackberry?s_trg=11&cd=1",
		}, {
			Title: "BQ",
			Href:  "/rossiya/telefony/bq?s_trg=11&cd=1",
		}, {
			Title: "DEXP",
			Href:  "/rossiya/telefony/dexp?s_trg=11&cd=1",
		}, {
			Title: "Explay",
			Href:  "/rossiya/telefony/explay?s_trg=11&cd=1",
		}, {
			Title: "Fly",
			Href:  "/rossiya/telefony/fly?s_trg=11&cd=1",
		}, {
			Title: "Highscreen",
			Href:  "/rossiya/telefony/highscreen?s_trg=11&cd=1",
		}, {
			Title: "HTC",
			Href:  "/rossiya/telefony/htc?s_trg=11&cd=1",
		}, {
			Title: "Huawei",
			Href:  "/rossiya/telefony/huawei?s_trg=11&cd=1",
		}, {
			Title: "iPhone",
			Href:  "/rossiya/telefony/iphone?s_trg=11&cd=1",
		}, {
			Title: "Lenovo",
			Href:  "/rossiya/telefony/lenovo?s_trg=11&cd=1",
		}, {
			Title: "LG",
			Href:  "/rossiya/telefony/lg?s_trg=11&cd=1",
		}, {
			Title: "Meizu",
			Href:  "/rossiya/telefony/meizu?s_trg=11&cd=1",
		}, {
			Title: "Micromax",
			Href:  "/rossiya/telefony/micromax?s_trg=11&cd=1",
		}, {
			Title: "Microsoft",
			Href:  "/rossiya/telefony/microsoft?s_trg=11&cd=1",
		}, {
			Title: "Motorola",
			Href:  "/rossiya/telefony/motorola?s_trg=11&cd=1",
		}, {
			Title: "MTS",
			Href:  "/rossiya/telefony/mts?s_trg=11&cd=1",
		}, {
			Title: "Nokia",
			Href:  "/rossiya/telefony/nokia?s_trg=11&cd=1",
		}, {
			Title: "Panasonic",
			Href:  "/rossiya/telefony/panasonic?s_trg=11&cd=1",
		}, {
			Title: "Philips",
			Href:  "/rossiya/telefony/philips?s_trg=11&cd=1",
		}, {
			Title: "Prestigio",
			Href:  "/rossiya/telefony/prestigio?s_trg=11&cd=1",
		}, {
			Title: "Samsung",
			Href:  "/rossiya/telefony/samsung?s_trg=11&cd=1",
		}, {
			Title: "Siemens",
			Href:  "/rossiya/telefony/siemens?s_trg=11&cd=1",
		}, {
			Title: "SkyLink",
			Href:  "/rossiya/telefony/skylink?s_trg=11&cd=1",
		}, {
			Title: "Sony",
			Href:  "/rossiya/telefony/sony?s_trg=11&cd=1",
		}, {
			Title: "teXet",
			Href:  "/rossiya/telefony/texet?s_trg=11&cd=1",
		}, {
			Title: "Vertu",
			Href:  "/rossiya/telefony/vertu?s_trg=11&cd=1",
		}, {
			Title: "Xiaomi",
			Href:  "/rossiya/telefony/xiaomi?s_trg=11&cd=1",
		}, {
			Title: "ZTE",
			Href:  "/rossiya/telefony/zte?s_trg=11&cd=1",
		}, {
			Title: "Другие марки",
			Href:  "/rossiya/telefony/drugie_marki?s_trg=11&cd=1",
		}, {
			Title: "Рации",
			Href:  "/rossiya/telefony/ratsii?s_trg=11&cd=1",
		}, {
			Title: "Стационарные телефоны",
			Href:  "/rossiya/telefony/statsionarnye_telefony?s_trg=11&cd=1",
		}, {
			Title: "Аксессуары",
			Href:  "/rossiya/telefony/aksessuary?s_trg=11&cd=1",
			Childes: []category_item.CategoryItem{
				{
					Title: "Аккумуляторы",
					Href:  "/rossiya/telefony/aksessuary/akkumulyatory?s_trg=11&cd=1",
				}, {
					Title: "Гарнитуры и наушники",
					Href:  "/rossiya/telefony/aksessuary/garnitury_i_naushniki?s_trg=11&cd=1",
				}, {
					Title: "Зарядные устройства",
					Href:  "/rossiya/telefony/aksessuary/zaryadnye_ustroystva?s_trg=11&cd=1",
				}, {
					Title: "Кабели и адаптеры",
					Href:  "/rossiya/telefony/aksessuary/kabeli_i_adaptery?s_trg=11&cd=1",
				}, {
					Title: "Модемы и роутеры",
					Href:  "/rossiya/telefony/aksessuary/modemy_i_routery?s_trg=11&cd=1",
				}, {
					Title: "Чехлы и плёнки",
					Href:  "/rossiya/telefony/aksessuary/chehly_i_plenki?s_trg=11&cd=1",
				}, {
					Title: "Запчасти",
					Href:  "/rossiya/telefony/aksessuary/zapchasti?s_trg=11&cd=1"}}}},
	}, {
		Title: "Товары для компьютера",
		Href:  "/rossiya/tovary_dlya_kompyutera?s_trg=11&cd=1",
		Childes: []category_item.CategoryItem{{
			Title: "Акустика",
			Href:  "/rossiya/tovary_dlya_kompyutera/akustika?s_trg=11&cd=1",
		}, {
			Title: "Веб-камеры",
			Href:  "/rossiya/tovary_dlya_kompyutera/veb-kamery?s_trg=11&cd=1",
		}, {
			Title: "Джойстики и рули",
			Href:  "/rossiya/tovary_dlya_kompyutera/dzhoystiki_i_ruli?s_trg=11&cd=1",
		}, {
			Title: "Клавиатуры и мыши",
			Href:  "/rossiya/tovary_dlya_kompyutera/klaviatury_i_myshi?s_trg=11&cd=1",
		}, {
			Title: "Комплектующие",
			Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie?s_trg=11&cd=1",
			Childes: []category_item.CategoryItem{{
				Title: "CD, DVD и Blu-ray приводы",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/cd_dvd_i_blu-ray_privody?s_trg=11&cd=1",
			}, {
				Title: "Блоки питания",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/bloki_pitaniya?s_trg=11&cd=1",
			}, {
				Title: "Видеокарты",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/videokarty?s_trg=11&cd=1",
			}, {
				Title: "Жёсткие диски",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/zhestkie_diski?s_trg=11&cd=1",
			}, {
				Title: "Звуковые карты",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/zvukovye_karty?s_trg=11&cd=1",
			}, {
				Title: "Контроллеры",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/kontrollery?s_trg=11&cd=1",
			}, {
				Title: "Корпусы",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/korpusy?s_trg=11&cd=1",
			}, {
				Title: "Материнские платы",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/materinskie_platy?s_trg=11&cd=1",
			}, {
				Title: "Оперативная память",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/operativnaya_pamyat?s_trg=11&cd=1",
			}, {
				Title: "Процессоры",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/protsessory?s_trg=11&cd=1",
			}, {
				Title: "Системы охлаждения",
				Href:  "/rossiya/tovary_dlya_kompyutera/komplektuyuschie/sistemy_ohlazhdeniya?s_trg=11&cd=1"}},
		}, {
			Title: "Мониторы",
			Href:  "/rossiya/tovary_dlya_kompyutera/monitory?s_trg=11&cd=1",
		}, {
			Title: "Переносные жёсткие диски",
			Href:  "/rossiya/tovary_dlya_kompyutera/perenosnye_zhestkie_diski?s_trg=11&cd=1",
		}, {
			Title: "Сетевое оборудование",
			Href:  "/rossiya/tovary_dlya_kompyutera/setevoe_oborudovanie?s_trg=11&cd=1",
		}, {
			Title: "ТВ-тюнеры",
			Href:  "/rossiya/tovary_dlya_kompyutera/tv-tyunery?s_trg=11&cd=1",
		}, {
			Title: "Флэшки и карты памяти",
			Href:  "/rossiya/tovary_dlya_kompyutera/fleshki_i_karty_pamyati?s_trg=11&cd=1",
		}, {
			Title: "Аксессуары",
			Href:  "/rossiya/tovary_dlya_kompyutera/aksessuary?s_trg=11&cd=1"}},
	}, {
		Title: "Фототехника",
		Href:  "/rossiya/fototehnika?s_trg=11&cd=1",
		Childes: []category_item.CategoryItem{{
			Title: "Компактные фотоаппараты",
			Href:  "/rossiya/fototehnika/kompaktnye_fotoapparaty?s_trg=11&cd=1",
		}, {
			Title: "Зеркальные фотоаппараты",
			Href:  "/rossiya/fototehnika/zerkalnye_fotoapparaty?s_trg=11&cd=1",
		}, {
			Title: "Плёночные фотоаппараты",
			Href:  "/rossiya/fototehnika/plenochnye_fotoapparaty?s_trg=11&cd=1",
		}, {
			Title: "Бинокли и телескопы",
			Href:  "/rossiya/fototehnika/binokli_i_teleskopy?s_trg=11&cd=1",
		}, {
			Title: "Объективы",
			Href:  "/rossiya/fototehnika/obektivy?s_trg=11&cd=1",
		}, {
			Title: "Оборудование и аксессуары",
			Href:  "/rossiya/fototehnika/oborudovanie_i_aksessuary?s_trg=11&cd=1"}}}}}}

func TestRecursiveSearch(t *testing.T) {
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(inputHtml))
	if err != nil {
		t.Fatal(err)
	}

	ul := document.Find("div[class^='rubricator-root'] > ul[class^='rubricator-list']").Nodes[0]

	result := recursiveParser(ul)

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("expected:\n\n%v\n\ngot:\n\n%v\n\n", expectedOutput, result)
	}
}
