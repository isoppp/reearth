# Changelog

All notable changes to this project will be documented in this file.

## 0.15.0 - 2023-05-18

### Web

#### 🚀 Features

- Support nested style expression object in style evaluator ([#423](https://github.com/reearth/reearth/pull/423)) [`60b52c`](https://github.com/reearth/reearth/commit/60b52c)
- Add reearth ver support in profile header ([#428](https://github.com/reearth/reearth/pull/428)) [`d1125e`](https://github.com/reearth/reearth/commit/d1125e)
- Sandbox plugin iframe - alpha ([#399](https://github.com/reearth/reearth/pull/399)) [`d77fcb`](https://github.com/reearth/reearth/commit/d77fcb)
- Add netlify to web ([#413](https://github.com/reearth/reearth/pull/413)) [`ce6935`](https://github.com/reearth/reearth/commit/ce6935)
- Support csv download for Dataset in setting page  ([#410](https://github.com/reearth/reearth/pull/410)) [`3bb0a0`](https://github.com/reearth/reearth/commit/3bb0a0)
- Properties for widget areas ([#427](https://github.com/reearth/reearth/pull/427)) [`e6e805`](https://github.com/reearth/reearth/commit/e6e805)

#### 🔧 Bug Fixes

- Add data url to componentId for better tracking of component in core ([#421](https://github.com/reearth/reearth/pull/421)) [`c3ab03`](https://github.com/reearth/reearth/commit/c3ab03)
- Update cesium version to 1.105.2 from 1.105.1 ([#425](https://github.com/reearth/reearth/pull/425)) [`de8896`](https://github.com/reearth/reearth/commit/de8896)
- Create new workspace doesn&[#39](https://github.com/reearth/reearth/pull/39);t save session ([#417](https://github.com/reearth/reearth/pull/417)) [`2c36e7`](https://github.com/reearth/reearth/commit/2c36e7)
- Classic infobox always shown for 3d tiles ([#416](https://github.com/reearth/reearth/pull/416)) [`a7a944`](https://github.com/reearth/reearth/commit/a7a944)
- Wrong last workspace value when multiple tabs ([#405](https://github.com/reearth/reearth/pull/405)) [`d68440`](https://github.com/reearth/reearth/commit/d68440)
- Feature containing layer show hide not working ([#412](https://github.com/reearth/reearth/pull/412)) [`2a97ef`](https://github.com/reearth/reearth/commit/2a97ef)
- Point cloud style converse error with undefined ([#398](https://github.com/reearth/reearth/pull/398)) [`9bf955`](https://github.com/reearth/reearth/commit/9bf955)
- Remove package.json yarn.lock from root ([#395](https://github.com/reearth/reearth/pull/395)) [`a87a2c`](https://github.com/reearth/reearth/commit/a87a2c)
- Pre-commit hook with husky [`b96641`](https://github.com/reearth/reearth/commit/b96641)
- Widget align system ([#436](https://github.com/reearth/reearth/pull/436)) [`0b8c85`](https://github.com/reearth/reearth/commit/0b8c85)

#### Miscellaneous Tasks

- Upgrade auth0 to version 2 ([#420](https://github.com/reearth/reearth/pull/420)) [`abfb4d`](https://github.com/reearth/reearth/commit/abfb4d)
- Update dependencies ([#411](https://github.com/reearth/reearth/pull/411)) [`58632a`](https://github.com/reearth/reearth/commit/58632a)
- Remove redundant comments [`57aa3c`](https://github.com/reearth/reearth/commit/57aa3c)
- Purge reearth-web of reference to team-Setting page ([#394](https://github.com/reearth/reearth/pull/394)) [`bc6d01`](https://github.com/reearth/reearth/commit/bc6d01)

### Server

#### 🚀 Features

- Export dataset as csv ([#409](https://github.com/reearth/reearth/pull/409)) [`79077b`](https://github.com/reearth/reearth/commit/79077b)
- Add experimental_sandbox option to scene settings [`8d688c`](https://github.com/reearth/reearth/commit/8d688c)
- REEARTH_WEB_CONFIG envvar [`e6e79e`](https://github.com/reearth/reearth/commit/e6e79e)
- Vr mode in scene property [`53030d`](https://github.com/reearth/reearth/commit/53030d)
- Add config to disable web feature [`654758`](https://github.com/reearth/reearth/commit/654758)
- Serve published pages on root path ([#386](https://github.com/reearth/reearth/pull/386)) [`845531`](https://github.com/reearth/reearth/commit/845531)
- Label background padding for markers [`160944`](https://github.com/reearth/reearth/commit/160944)
- Label background color property field for markers [`ac10b4`](https://github.com/reearth/reearth/commit/ac10b4)
- Add padding properties to widget align system area ([#381](https://github.com/reearth/reearth/pull/381)) [`d5dbcf`](https://github.com/reearth/reearth/commit/d5dbcf)
- Add Visible field to built-in widgets ([#380](https://github.com/reearth/reearth/pull/380)) [`8c1d82`](https://github.com/reearth/reearth/commit/8c1d82)

#### 🔧 Bug Fixes

- Fix favicon url [`12dafb`](https://github.com/reearth/reearth/commit/12dafb)
- Rewrite title and favicon for published html [`03b5be`](https://github.com/reearth/reearth/commit/03b5be)
- Env vars to change html title and favicon ([#390](https://github.com/reearth/reearth/pull/390)) [`a9910c`](https://github.com/reearth/reearth/commit/a9910c)
- Upload assets outside transaction [`afc7e2`](https://github.com/reearth/reearth/commit/afc7e2)
- Handle transaction correctly [`4b164b`](https://github.com/reearth/reearth/commit/4b164b)
- Prefer dataset values when merging properties ([#388](https://github.com/reearth/reearth/pull/388)) [`8112f7`](https://github.com/reearth/reearth/commit/8112f7)
- Add schema to published field in web config [`b17807`](https://github.com/reearth/reearth/commit/b17807)
- Add published field to web config [`18fdf7`](https://github.com/reearth/reearth/commit/18fdf7)
- Description of experimental flag in scene property [`b04157`](https://github.com/reearth/reearth/commit/b04157)
- Htmlblock order ([#385](https://github.com/reearth/reearth/pull/385)) [`bb1b9c`](https://github.com/reearth/reearth/commit/bb1b9c)
- Use auth0 web client id for reearth_config.json [`634799`](https://github.com/reearth/reearth/commit/634799)
- Make widget area gap optional ([#383](https://github.com/reearth/reearth/pull/383)) [`1e00ca`](https://github.com/reearth/reearth/commit/1e00ca)
- Widget padding fixes ([#382](https://github.com/reearth/reearth/pull/382)) [`a019c5`](https://github.com/reearth/reearth/commit/a019c5)
- Widget area could not be saved and loaded to mongo [`3f3adf`](https://github.com/reearth/reearth/commit/3f3adf)
- Public plugin was deleted on installing privte plugin that has same name [`d24034`](https://github.com/reearth/reearth/commit/d24034)

#### ✨ Refactor

- Use reearthx cache ([#401](https://github.com/reearth/reearth/pull/401)) [`c4811d`](https://github.com/reearth/reearth/commit/c4811d)

#### 🧪 Testing

- Some unit test fails on windows ([#392](https://github.com/reearth/reearth/pull/392)) [`bc89f3`](https://github.com/reearth/reearth/commit/bc89f3)

#### Miscellaneous Tasks

- Lowercase the error message [`c22565`](https://github.com/reearth/reearth/commit/c22565)
- Upgrade ci go linter to 1.52.* ([#402](https://github.com/reearth/reearth/pull/402)) [`2a87bb`](https://github.com/reearth/reearth/commit/2a87bb)
- Advanced option to scene settings [`707601`](https://github.com/reearth/reearth/commit/707601)
- Upgrade deps [`bed3f2`](https://github.com/reearth/reearth/commit/bed3f2)
- Add logs to plugin download from marketplace [`436b68`](https://github.com/reearth/reearth/commit/436b68)

### Misc

#### 🚀 Features

- Support polyline and point on reearth&#x2F;core ([#606](https://github.com/reearth/reearth/pull/606)) [`abd9a3`](https://github.com/reearth/reearth/commit/abd9a3)
- Add remembering last open workspace functionality ([#598](https://github.com/reearth/reearth/pull/598)) [`968a9d`](https://github.com/reearth/reearth/commit/968a9d)
- Support TileMapService(TMS) on reearth&#x2F;core ([#604](https://github.com/reearth/reearth/pull/604)) [`284b35`](https://github.com/reearth/reearth/commit/284b35)
- Support styling color &amp; show for 3dtiles model ([#599](https://github.com/reearth/reearth/pull/599)) [`a82ca2`](https://github.com/reearth/reearth/commit/a82ca2)
- Support tiles data type on reearth&#x2F;core ([#597](https://github.com/reearth/reearth/pull/597) [`ae5c49`](https://github.com/reearth/reearth/commit/ae5c49)
- Support classificationType property in some feature on reearth&#x2F;core ([#593](https://github.com/reearth/reearth/pull/593)) [`897868`](https://github.com/reearth/reearth/commit/897868)
- Upgrade cesium-mvt-imagery-provider ([#591](https://github.com/reearth/reearth/pull/591)) [`d01d01`](https://github.com/reearth/reearth/commit/d01d01)
- Upgrade cesium-mvt-imagery-provider ([#588](https://github.com/reearth/reearth/pull/588)) [`e138bd`](https://github.com/reearth/reearth/commit/e138bd)
- Add more styling properties to resource appearance in reearth&#x2F;core ([#586](https://github.com/reearth/reearth/pull/586)) [`8f3625`](https://github.com/reearth/reearth/commit/8f3625)
- Add built field in scene in plugin api ([#584](https://github.com/reearth/reearth/pull/584)) [`e8050c`](https://github.com/reearth/reearth/commit/e8050c)
- Add scene light ([#576](https://github.com/reearth/reearth/pull/576)) [`43c2b8`](https://github.com/reearth/reearth/commit/43c2b8)
- Upgrade mvt lib ([#575](https://github.com/reearth/reearth/pull/575)) [`584cee`](https://github.com/reearth/reearth/commit/584cee)
- Option to unselect layer when click infobox close ([#564](https://github.com/reearth/reearth/pull/564)) [`f2b2f2`](https://github.com/reearth/reearth/commit/f2b2f2)
- Add alpha property to raster appearance on reearth&#x2F;core ([#555](https://github.com/reearth/reearth/pull/555)) [`541f30`](https://github.com/reearth/reearth/commit/541f30)
- Support styling for point cloud on reearth&#x2F;core ([#549](https://github.com/reearth/reearth/pull/549)) [`f410d6`](https://github.com/reearth/reearth/commit/f410d6)
- Add defaultContent property for infobox in plugin API on reearth&#x2F;core ([#538](https://github.com/reearth/reearth/pull/538)) [`31ba31`](https://github.com/reearth/reearth/commit/31ba31)
- Add updateClockOnLoad  to data on reearth&#x2F;core ([#539](https://github.com/reearth/reearth/pull/539)) [`3653e2`](https://github.com/reearth/reearth/commit/3653e2)
- Support gltf data type on reearth&#x2F;core ([#535](https://github.com/reearth/reearth/pull/535)) [`e086be`](https://github.com/reearth/reearth/commit/e086be)
- Support ga4 ([#509](https://github.com/reearth/reearth/pull/509)) [`39bfb0`](https://github.com/reearth/reearth/commit/39bfb0)
- Select MVT feature on reearth&#x2F;core ([#527](https://github.com/reearth/reearth/pull/527)) [`de605d`](https://github.com/reearth/reearth/commit/de605d)
- Add parameters property to data on reearth&#x2F;core ([#520](https://github.com/reearth/reearth/pull/520)) [`c698eb`](https://github.com/reearth/reearth/commit/c698eb)
- Layer select event in reearth&#x2F;core ([#470](https://github.com/reearth/reearth/pull/470)) [`fb22e6`](https://github.com/reearth/reearth/commit/fb22e6)
- Show field in appearances of reearth&#x2F;core ([#469](https://github.com/reearth/reearth/pull/469)) [`819eb6`](https://github.com/reearth/reearth/commit/819eb6)
- ImageSizeInMeters field in marker proeperty ([#511](https://github.com/reearth/reearth/pull/511)) [`290cb7`](https://github.com/reearth/reearth/commit/290cb7)
- Override clock from scene setting on reearth&#x2F;core ([#505](https://github.com/reearth/reearth/pull/505)) [`01bffd`](https://github.com/reearth/reearth/commit/01bffd)
- Support features for CZML on reearth&#x2F;core ([#506](https://github.com/reearth/reearth/pull/506)) [`e5c160`](https://github.com/reearth/reearth/commit/e5c160)
- Update cesium ([#503](https://github.com/reearth/reearth/pull/503)) [`5a649f`](https://github.com/reearth/reearth/commit/5a649f)
- Color blend mode in tileset on reearth&#x2F;core ([#496](https://github.com/reearth/reearth/pull/496)) [`ca43dc`](https://github.com/reearth/reearth/commit/ca43dc)
- Change brand images and colors at the root page ([#495](https://github.com/reearth/reearth/pull/495))Co-authored-by: rot1024 &lt;aayhrot@gmail.com&gt; [`4f07b9`](https://github.com/reearth/reearth/commit/4f07b9)
- Add builtin clipping box on reearth&#x2F;core ([#487](https://github.com/reearth/reearth/pull/487)) [`63bd4f`](https://github.com/reearth/reearth/commit/63bd4f)
- ExtrudedHeight for polygon on reearth&#x2F;core ([#486](https://github.com/reearth/reearth/pull/486)) [`523d35`](https://github.com/reearth/reearth/commit/523d35)
- Support resource entity layerId on reearth&#x2F;core ([#485](https://github.com/reearth/reearth/pull/485)) [`7bd7c5`](https://github.com/reearth/reearth/commit/7bd7c5)
- Support distanceDisplayCondition on reearth&#x2F;core ([#475](https://github.com/reearth/reearth/pull/475)) [`ce8270`](https://github.com/reearth/reearth/commit/ce8270)
- Support GeoRSS and gml in reearth&#x2F;core ([#455](https://github.com/reearth/reearth/pull/455)) [`58c25b`](https://github.com/reearth/reearth/commit/58c25b)
- Add htmlblock on reearth&#x2F;core ([#454](https://github.com/reearth/reearth/pull/454)) [`1b37e0`](https://github.com/reearth/reearth/commit/1b37e0)
- Support clipping box direction on reearth&#x2F;core ([#467](https://github.com/reearth/reearth/pull/467)) [`70f74e`](https://github.com/reearth/reearth/commit/70f74e)
- Add sampleTerrainHeight on reearht&#x2F;core ([#466](https://github.com/reearth/reearth/pull/466)) [`55334e`](https://github.com/reearth/reearth/commit/55334e)
- Get brand from config ([#457](https://github.com/reearth/reearth/pull/457)) [`d35361`](https://github.com/reearth/reearth/commit/d35361)
- Support timeline on mobile on reearth&#x2F;core ([#462](https://github.com/reearth/reearth/pull/462)) [`efeaf4`](https://github.com/reearth/reearth/commit/efeaf4)
- Use visible field on reearth&#x2F;core ([#456](https://github.com/reearth/reearth/pull/456)) [`333610`](https://github.com/reearth/reearth/commit/333610)
- Add htmlblock to built-in plugin ([#384](https://github.com/reearth/reearth/pull/384)) [`51c79a`](https://github.com/reearth/reearth/commit/51c79a)
- Add override, replace and delete plugin API on reearth&#x2F;core ([#451](https://github.com/reearth/reearth/pull/451)) [`2e1c41`](https://github.com/reearth/reearth/commit/2e1c41)
- Selecting imagery features ([#450](https://github.com/reearth/reearth/pull/450)) [`f24ef5`](https://github.com/reearth/reearth/commit/f24ef5)
- Support interval fetching data on reearth&#x2F;core ([#449](https://github.com/reearth/reearth/pull/449)) [`406174`](https://github.com/reearth/reearth/commit/406174)
- Support select feature on reearth&#x2F;core ([#445](https://github.com/reearth/reearth/pull/445)) [`3174b1`](https://github.com/reearth/reearth/commit/3174b1)
- Use experimental core flag ([#448](https://github.com/reearth/reearth/pull/448)) [`b04294`](https://github.com/reearth/reearth/commit/b04294)
- Support time series features on reearth&#x2F;core ([#446](https://github.com/reearth/reearth/pull/446)) [`8fc9b6`](https://github.com/reearth/reearth/commit/8fc9b6)
- Scene property to enable VR mode ([#444](https://github.com/reearth/reearth/pull/444)) [`3d35aa`](https://github.com/reearth/reearth/commit/3d35aa)
- Support general transit feed ([#408](https://github.com/reearth/reearth/pull/408)) [`49b4a4`](https://github.com/reearth/reearth/commit/49b4a4)
- Support osm data type on reearth&#x2F;core ([#431](https://github.com/reearth/reearth/pull/431)) [`0d4e0b`](https://github.com/reearth/reearth/commit/0d4e0b)
- Label background color and padding property for markers ([#426](https://github.com/reearth/reearth/pull/426) [`72cd0d`](https://github.com/reearth/reearth/commit/72cd0d)
- Support entity base flyTo on reearth&#x2F;core ([#419](https://github.com/reearth/reearth/pull/419)) [`3060cf`](https://github.com/reearth/reearth/commit/3060cf)
- Support overriding czml appearance on reearth&#x2F;core ([#421](https://github.com/reearth/reearth/pull/421)) [`e62f4d`](https://github.com/reearth/reearth/commit/e62f4d)
- Support kml on reearth&#x2F;core ([#422](https://github.com/reearth/reearth/pull/422)) [`052daf`](https://github.com/reearth/reearth/commit/052daf)
- Support json properties on reearth&#x2F;core ([#412](https://github.com/reearth/reearth/pull/412)) [`ac7986`](https://github.com/reearth/reearth/commit/ac7986)
- Connect reearth&#x2F;core with existence pages ([#401](https://github.com/reearth/reearth/pull/401)) [`0735c0`](https://github.com/reearth/reearth/commit/0735c0)
- Add runTimes property to PluginInstance ([#404](https://github.com/reearth/reearth/pull/404)) [`17d787`](https://github.com/reearth/reearth/commit/17d787)
- Support plugin system on reearth&#x2F;core ([#399](https://github.com/reearth/reearth/pull/399)) [`bab9e6`](https://github.com/reearth/reearth/commit/bab9e6)
- Add selectedFeature and selectedComputedFeature on reearth&#x2F;core ([#398](https://github.com/reearth/reearth/pull/398)) [`474b34`](https://github.com/reearth/reearth/commit/474b34)
- Set modal above popup ([#397](https://github.com/reearth/reearth/pull/397)) [`ff47c5`](https://github.com/reearth/reearth/commit/ff47c5)
- Port 2d navigator to reearth&#x2F;core ([#394](https://github.com/reearth/reearth/pull/394)) [`07a6b4`](https://github.com/reearth/reearth/commit/07a6b4)
- Core&#x2F;Visualizer without plugins ([#372](https://github.com/reearth/reearth/pull/372)) [`f97c38`](https://github.com/reearth/reearth/commit/f97c38)
- Reearth style language ([#384](https://github.com/reearth/reearth/pull/384)) [`a828ac`](https://github.com/reearth/reearth/commit/a828ac)
- Support 3dtiles on reearth&#x2F;core ([#392](https://github.com/reearth/reearth/pull/392)) [`e8068f`](https://github.com/reearth/reearth/commit/e8068f)
- Support MVT on reearth&#x2F;core ([#388](https://github.com/reearth/reearth/pull/388)) [`cac89c`](https://github.com/reearth/reearth/commit/cac89c)
- Support WMS on reearth&#x2F;core ([#387](https://github.com/reearth/reearth/pull/387)) [`666c1b`](https://github.com/reearth/reearth/commit/666c1b)
- Support CZML in reearth&#x2F;core ([#383](https://github.com/reearth/reearth/pull/383)) [`f44d98`](https://github.com/reearth/reearth/commit/f44d98)
- Plugin api client storage ([#376](https://github.com/reearth/reearth/pull/376)) [`4f36ad`](https://github.com/reearth/reearth/commit/4f36ad)
- Support csv on the reearth&#x2F;core ([#382](https://github.com/reearth/reearth/pull/382)) [`a8f5bf`](https://github.com/reearth/reearth/commit/a8f5bf)

#### 🔧 Bug Fixes

- Merge conflict resolved [`39a37a`](https://github.com/reearth/reearth/commit/39a37a)
- Use core hook may get value incorrectly ([#609](https://github.com/reearth/reearth/pull/609)) [`d87e90`](https://github.com/reearth/reearth/commit/d87e90)
- Resolve endsWith is not a function error with style in reearth&#x2F;core ([#603](https://github.com/reearth/reearth/pull/603)) [`15fec3`](https://github.com/reearth/reearth/commit/15fec3)
- Add % as special case for &#x60;Number()&#x60; in style lang of reearth&#x2F;core ([#601](https://github.com/reearth/reearth/pull/601)) [`5a4a3d`](https://github.com/reearth/reearth/commit/5a4a3d)
- Rename moveWidget to onMoveWidget on reearth&#x2F;core ([#600](https://github.com/reearth/reearth/pull/600)) [`6c06ab`](https://github.com/reearth/reearth/commit/6c06ab)
- Enable splash screen in preview page ([#596](https://github.com/reearth/reearth/pull/596) [`e1f5ac`](https://github.com/reearth/reearth/commit/e1f5ac)
- Timeline scroll should be fixed in initial render ([#571](https://github.com/reearth/reearth/pull/571)) [`6a0aed`](https://github.com/reearth/reearth/commit/6a0aed)
- Remove default height reference for modelGraphics in reearth&#x2F;core ([#592](https://github.com/reearth/reearth/pull/592)) [`961b46`](https://github.com/reearth/reearth/commit/961b46)
- Remove duplication of feature entity in reearth&#x2F;core ([#590](https://github.com/reearth/reearth/pull/590)) [`8cc03e`](https://github.com/reearth/reearth/commit/8cc03e)
- Prevent selecting not shown feature for MVT on reearth&#x2F;core ([#589](https://github.com/reearth/reearth/pull/589)) [`41816d`](https://github.com/reearth/reearth/commit/41816d)
- Infobox for 3dtiles feature ([#587](https://github.com/reearth/reearth/pull/587)) [`70cfdb`](https://github.com/reearth/reearth/commit/70cfdb)
- Use layer id with feature id for entity id on reearth&#x2F;core ([#585](https://github.com/reearth/reearth/pull/585)) [`24cc88`](https://github.com/reearth/reearth/commit/24cc88)
- Can&[#39](https://github.com/reearth/reearth/pull/39);t set iframe&[#39](https://github.com/reearth/reearth/pull/39);s width or height individually in crust  ([#583](https://github.com/reearth/reearth/pull/583)) [`3494b4`](https://github.com/reearth/reearth/commit/3494b4)
- MVT isn&[#39](https://github.com/reearth/reearth/pull/39);t render correctly on reearth&#x2F;core ([#582](https://github.com/reearth/reearth/pull/582)) [`988d1a`](https://github.com/reearth/reearth/commit/988d1a)
- Sample terrain height API cannot properly return promise ([#581](https://github.com/reearth/reearth/pull/581)) [`72eafd`](https://github.com/reearth/reearth/commit/72eafd)
- Change logical operator evaluation behaviour in reearth&#x2F;core ([#580](https://github.com/reearth/reearth/pull/580)) [`3d23ce`](https://github.com/reearth/reearth/commit/3d23ce)
- Infobox style collapse for long names ([#578](https://github.com/reearth/reearth/pull/578)) [`04775d`](https://github.com/reearth/reearth/commit/04775d)
- Override currentTime on timeline widget when time is updated by CZML on reearth&#x2F;core ([#579](https://github.com/reearth/reearth/pull/579)) [`dbe2e1`](https://github.com/reearth/reearth/commit/dbe2e1)
- Rename scene light properties [`8b3a18`](https://github.com/reearth/reearth/commit/8b3a18)
- Add none cesium value for heightReference ([#577](https://github.com/reearth/reearth/pull/577)) [`2c92e3`](https://github.com/reearth/reearth/commit/2c92e3)
- JSON Path condition on reearth&#x2F;core ([#572](https://github.com/reearth/reearth/pull/572)) [`5f7024`](https://github.com/reearth/reearth/commit/5f7024)
- Update html head [`f1780f`](https://github.com/reearth/reearth/commit/f1780f)
- Sync selected feature with layer api on reearth&#x2F;core ([#570](https://github.com/reearth/reearth/pull/570)) [`ded6b0`](https://github.com/reearth/reearth/commit/ded6b0)
- Imagery index on reearth&#x2F;core ([#569](https://github.com/reearth/reearth/pull/569)) [`6b233f`](https://github.com/reearth/reearth/commit/6b233f)
- Revert add height to polygon in reearth&#x2F;core ([#566](https://github.com/reearth/reearth/pull/566)) [`1f2f74`](https://github.com/reearth/reearth/commit/1f2f74)
- Add clamp as default height reference for model in reearth&#x2F;core ([#567](https://github.com/reearth/reearth/pull/567)) [`517386`](https://github.com/reearth/reearth/commit/517386)
- Minimum timeline range on reearth&#x2F;core ([#565](https://github.com/reearth/reearth/pull/565)) [`975c79`](https://github.com/reearth/reearth/commit/975c79)
- Infobox html block styling ([#562](https://github.com/reearth/reearth/pull/562)) [`32b248`](https://github.com/reearth/reearth/commit/32b248)
- Imagery layer tile index on reearth&#x2F;core ([#561](https://github.com/reearth/reearth/pull/561) [`25bdff`](https://github.com/reearth/reearth/commit/25bdff)
- Add &quot;disabled&quot; as shadow mode on reearth&#x2F;core ([#560](https://github.com/reearth/reearth/pull/560)) [`1632ea`](https://github.com/reearth/reearth/commit/1632ea)
- Support multi layers in MVT on reearth&#x2F;core ([#559](https://github.com/reearth/reearth/pull/559)) [`9c8c88`](https://github.com/reearth/reearth/commit/9c8c88)
- Html block margin and height [`09f6ef`](https://github.com/reearth/reearth/commit/09f6ef)
- UpdateClockOnLoad condition on reearth&#x2F;core ([#558](https://github.com/reearth/reearth/pull/558)) [`f9076a`](https://github.com/reearth/reearth/commit/f9076a)
- Add height to polygon in reearth&#x2F;core ([#557](https://github.com/reearth/reearth/pull/557)) [`a84dbe`](https://github.com/reearth/reearth/commit/a84dbe)
- Use computed feature on resource on reearth&#x2F;core ([#556](https://github.com/reearth/reearth/pull/556)) [`006f40`](https://github.com/reearth/reearth/commit/006f40)
- Point cloud dirty check on reearth&#x2F;core ([#554](https://github.com/reearth/reearth/pull/554)) [`f95d27`](https://github.com/reearth/reearth/commit/f95d27)
- Support default infobox and selection indicator on imagery layer on reearth&#x2F;core ([#553](https://github.com/reearth/reearth/pull/553)) [`1946b3`](https://github.com/reearth/reearth/commit/1946b3)
- Scrollbar in timeline widget always showing ([#550](https://github.com/reearth/reearth/pull/550)) [`84d63d`](https://github.com/reearth/reearth/commit/84d63d)
- Clipping box on point cloud on reearth&#x2F;core ([#552](https://github.com/reearth/reearth/pull/552)) [`fb4455`](https://github.com/reearth/reearth/commit/fb4455)
- Htmlblock on safari ([#548](https://github.com/reearth/reearth/pull/548)) [`392fc7`](https://github.com/reearth/reearth/commit/392fc7)
- Use attributes as default content for infobox on reearth&#x2F;core ([#547](https://github.com/reearth/reearth/pull/547)) [`be3718`](https://github.com/reearth/reearth/commit/be3718)
- Deleting feature process on reearth&#x2F;core ([#546](https://github.com/reearth/reearth/pull/546)) [`39ecaf`](https://github.com/reearth/reearth/commit/39ecaf)
- Trigger select event when featureId is changed on reearth&#x2F;core ([#545](https://github.com/reearth/reearth/pull/545)) [`6f5401`](https://github.com/reearth/reearth/commit/6f5401)
- Draw polylines on polygon on reearth&#x2F;core ([#544](https://github.com/reearth/reearth/pull/544)) [`d35329`](https://github.com/reearth/reearth/commit/d35329)
- Infobox html color ([#534](https://github.com/reearth/reearth/pull/534)) [`c0e0a6`](https://github.com/reearth/reearth/commit/c0e0a6)
- Use overriddenLayers to get infobox on reearth&#x2F;core ([#541](https://github.com/reearth/reearth/pull/541)) [`c4e9db`](https://github.com/reearth/reearth/commit/c4e9db)
- Fly to multiple entities added by a layer on reearth&#x2F;core ([#540](https://github.com/reearth/reearth/pull/540)) [`95b3d3`](https://github.com/reearth/reearth/commit/95b3d3)
- Add polyfill for requestIdleCallback in reearth&#x2F;core ([#537](https://github.com/reearth/reearth/pull/537)) [`c4722f`](https://github.com/reearth/reearth/commit/c4722f)
- Overriding timeline behavior on reearth&#x2F;core ([#532](https://github.com/reearth/reearth/pull/532)) [`890dae`](https://github.com/reearth/reearth/commit/890dae)
- Timeline bug on reearth&#x2F;core ([#531](https://github.com/reearth/reearth/pull/531)) [`572678`](https://github.com/reearth/reearth/commit/572678)
- Some error on reearth&#x2F;core ([#530](https://github.com/reearth/reearth/pull/530)) [`4ee0b6`](https://github.com/reearth/reearth/commit/4ee0b6)
- Abort fetching on data atom on reearth&#x2F;core ([#529](https://github.com/reearth/reearth/pull/529)) [`e82c88`](https://github.com/reearth/reearth/commit/e82c88)
- Timeline behavior on reearth&#x2F;core ([#528](https://github.com/reearth/reearth/pull/528)) [`461d06`](https://github.com/reearth/reearth/commit/461d06)
- Parse csv numeric strings as numbers in reearth&#x2F;core ([#526](https://github.com/reearth/reearth/pull/526)) [`9c890f`](https://github.com/reearth/reearth/commit/9c890f)
- Parse hyphen as reserved word in property key on reearth&#x2F;core ([#525](https://github.com/reearth/reearth/pull/525)) [`d49058`](https://github.com/reearth/reearth/commit/d49058)
- Condition for CZML on reearth&#x2F;core ([#524](https://github.com/reearth/reearth/pull/524)) [`3c61ca`](https://github.com/reearth/reearth/commit/3c61ca)
- Czml style for marker on reearth&#x2F;core ([#523](https://github.com/reearth/reearth/pull/523)) [`67dccb`](https://github.com/reearth/reearth/commit/67dccb)
- Use default block for entity on reearth&#x2F;core ([#522](https://github.com/reearth/reearth/pull/522)) [`ecd09a`](https://github.com/reearth/reearth/commit/ecd09a)
- Remove copyLazyLayer on reearth&#x2F;core ([#519](https://github.com/reearth/reearth/pull/519)) [`e17218`](https://github.com/reearth/reearth/commit/e17218)
- Copying lazy layers undefined behavior on reearth&#x2F;core ([#516](https://github.com/reearth/reearth/pull/516)) [`7b2c5e`](https://github.com/reearth/reearth/commit/7b2c5e)
- Copy lazy layer on plugin on reearth&#x2F;core ([#515](https://github.com/reearth/reearth/pull/515)) [`046a3d`](https://github.com/reearth/reearth/commit/046a3d)
- Attach style dynamically in resource on reearth&#x2F;core ([#514](https://github.com/reearth/reearth/pull/514)) [`81d291`](https://github.com/reearth/reearth/commit/81d291)
- Infinite infobox in CZML on reearth&#x2F;core ([#513](https://github.com/reearth/reearth/pull/513)) [`3e49a1`](https://github.com/reearth/reearth/commit/3e49a1)
- Selecting resource feature behavior on reearth&#x2F;core ([#512](https://github.com/reearth/reearth/pull/512)) [`35f2c2`](https://github.com/reearth/reearth/commit/35f2c2)
- Prevent unnecessary render on timeline on reearth&#x2F;core ([#510](https://github.com/reearth/reearth/pull/510)) [`7cb9d5`](https://github.com/reearth/reearth/commit/7cb9d5)
- Parse reserved word when property name includes reserved word on reearth&#x2F;core ([#508](https://github.com/reearth/reearth/pull/508)) [`aa247e`](https://github.com/reearth/reearth/commit/aa247e)
- Pass engine meta on plugin editor ([#507](https://github.com/reearth/reearth/pull/507)) [`9cd1b5`](https://github.com/reearth/reearth/commit/9cd1b5)
- Mvt line width on reearth&#x2F;core ([#504](https://github.com/reearth/reearth/pull/504)) [`c1a939`](https://github.com/reearth/reearth/commit/c1a939)
- It should not render entity when coordinate is undefined on reearth&#x2F;core [`23b6c6`](https://github.com/reearth/reearth/commit/23b6c6)
- Allow enter ground option for clipping box on reearth&#x2F;core ([#500](https://github.com/reearth/reearth/pull/500)) [`f8e129`](https://github.com/reearth/reearth/commit/f8e129)
- Clip area with clipping box on reearth&#x2F;core ([#498](https://github.com/reearth/reearth/pull/498)) [`4f647f`](https://github.com/reearth/reearth/commit/4f647f)
- Mvt cache on reearth&#x2F;core ([#497](https://github.com/reearth/reearth/pull/497)) [`ba2c7e`](https://github.com/reearth/reearth/commit/ba2c7e)
- Recreate no feature component when data url is changed on reearth&#x2F;core ([#494](https://github.com/reearth/reearth/pull/494)) [`93c805`](https://github.com/reearth/reearth/commit/93c805)
- Ignore cesium ion token when it is empty [`6648be`](https://github.com/reearth/reearth/commit/6648be)
- Feature type fix gpx on reearth&#x2F;core ([#493](https://github.com/reearth/reearth/pull/493)) [`51b8e8`](https://github.com/reearth/reearth/commit/51b8e8)
- Infobox error after layer delete ([#492](https://github.com/reearth/reearth/pull/492)) [`8a59dd`](https://github.com/reearth/reearth/commit/8a59dd)
- Delete bug that deletes necessary layer on reearth&#x2F;core ([#484](https://github.com/reearth/reearth/pull/484)) [`a0f48f`](https://github.com/reearth/reearth/commit/a0f48f)
- Force update when some data properties are updated on reearth&#x2F;core ([#483](https://github.com/reearth/reearth/pull/483)) [`867238`](https://github.com/reearth/reearth/commit/867238)
- Revert appearance to initial value when appearance is undefined on reearth&#x2F;core ([#482](https://github.com/reearth/reearth/pull/482)) [`76b6f4`](https://github.com/reearth/reearth/commit/76b6f4)
- Entity id is duplicated error on reearth&#x2F;core ([#481](https://github.com/reearth/reearth/pull/481)) [`a2b0b6`](https://github.com/reearth/reearth/commit/a2b0b6)
- Infobox property is undefined error on reearth&#x2F;core ([#480](https://github.com/reearth/reearth/pull/480)) [`2b8b07`](https://github.com/reearth/reearth/commit/2b8b07)
- Layers override behavior on rearth&#x2F;core ([#479](https://github.com/reearth/reearth/pull/479)) [`faff37`](https://github.com/reearth/reearth/commit/faff37)
- Error handling for time interval on reearth&#x2F;core ([#478](https://github.com/reearth/reearth/pull/478)) [`edf546`](https://github.com/reearth/reearth/commit/edf546)
- Overridden layers api on reearth&#x2F;core ([#477](https://github.com/reearth/reearth/pull/477)) [`abaade`](https://github.com/reearth/reearth/commit/abaade)
- Add properties in vehicle point for gtfs ([#476](https://github.com/reearth/reearth/pull/476)) [`1058ab`](https://github.com/reearth/reearth/commit/1058ab)
- Color function on reearth&#x2F;core ([#474](https://github.com/reearth/reearth/pull/474)) [`7d9fca`](https://github.com/reearth/reearth/commit/7d9fca)
- Errors when many functions are created in plugins ([#471](https://github.com/reearth/reearth/pull/471)) [`ebb50d`](https://github.com/reearth/reearth/commit/ebb50d)
- Coordinates for csv on reearth&#x2F;core ([#472](https://github.com/reearth/reearth/pull/472)) [`4a6473`](https://github.com/reearth/reearth/commit/4a6473)
- Lint and type error [`d00b9b`](https://github.com/reearth/reearth/commit/d00b9b)
- Suppress screen flicker when judging useCore [`b1852d`](https://github.com/reearth/reearth/commit/b1852d)
- Widget area has margin even if no widgets, disable widget area transition in built scene [`b51569`](https://github.com/reearth/reearth/commit/b51569)
- Cesium crashes when VR mode is false [`efa3fd`](https://github.com/reearth/reearth/commit/efa3fd)
- 3dtiles overriding appearance behavior on reearth&#x2F;core ([#468](https://github.com/reearth/reearth/pull/468)) [`8c48bc`](https://github.com/reearth/reearth/commit/8c48bc)
- Support visible and dynamic plane in clipping box on reearth&#x2F;core ([#465](https://github.com/reearth/reearth/pull/465)) [`4c89aa`](https://github.com/reearth/reearth/commit/4c89aa)
- SelectedFeature for 3dtiles on reearth&#x2F;core ([#463](https://github.com/reearth/reearth/pull/463)) [`cd1777`](https://github.com/reearth/reearth/commit/cd1777)
- Express undefined visible field on reearth&#x2F;core ([#461](https://github.com/reearth/reearth/pull/461)) [`c74630`](https://github.com/reearth/reearth/commit/c74630)
- Undefined behavior for visible field ([#460](https://github.com/reearth/reearth/pull/460)) [`c41d70`](https://github.com/reearth/reearth/commit/c41d70)
- Select entity on reearth&#x2F;core ([#458](https://github.com/reearth/reearth/pull/458)) [`bc1824`](https://github.com/reearth/reearth/commit/bc1824)
- Use default infobox on reearth&#x2F;core ([#453](https://github.com/reearth/reearth/pull/453)) [`d3fec8`](https://github.com/reearth/reearth/commit/d3fec8)
- Select event behavior on reearth&#x2F;core ([#452](https://github.com/reearth/reearth/pull/452)) [`384488`](https://github.com/reearth/reearth/commit/384488)
- Error in published page on reearth&#x2F;core ([#447](https://github.com/reearth/reearth/pull/447)) [`4c8805`](https://github.com/reearth/reearth/commit/4c8805)
- Expand timeline in initial load ([#443](https://github.com/reearth/reearth/pull/443)) [`d6a742`](https://github.com/reearth/reearth/commit/d6a742)
- Replace globe image when cesium ion token is updated ([#442](https://github.com/reearth/reearth/pull/442)) [`64ffae`](https://github.com/reearth/reearth/commit/64ffae)
- Layer fetch on reearth&#x2F;core ([#441](https://github.com/reearth/reearth/pull/441)) [`597b82`](https://github.com/reearth/reearth/commit/597b82)
- Dnd layer on reearth&#x2F;core ([#440](https://github.com/reearth/reearth/pull/440)) [`a5a2b4`](https://github.com/reearth/reearth/commit/a5a2b4)
- Disable requestRenderMode depends on widget on reearth&#x2F;core ([#439](https://github.com/reearth/reearth/pull/439)) [`12ce63`](https://github.com/reearth/reearth/commit/12ce63)
- Selected layer id is not propagated on reearth&#x2F;core ([#438](https://github.com/reearth/reearth/pull/438)) [`24993b`](https://github.com/reearth/reearth/commit/24993b)
- Handle featureId for 3dtiles and compat select plugin api on reearth&#x2F;core ([#417](https://github.com/reearth/reearth/pull/417)) [`9144ad`](https://github.com/reearth/reearth/commit/9144ad)
- Undefined behavior for resource on reearth&#x2F;core ([#437](https://github.com/reearth/reearth/pull/437)) [`3f51f2`](https://github.com/reearth/reearth/commit/3f51f2)
- Blocks cannot be displayed and updated as expected on reearth&#x2F;core ([#434](https://github.com/reearth/reearth/pull/434)) [`b5f921`](https://github.com/reearth/reearth/commit/b5f921)
- Layer appearances are not evaluated as expected ([#418](https://github.com/reearth/reearth/pull/418)) [`20382c`](https://github.com/reearth/reearth/commit/20382c)
- Support resource auto on reearth&#x2F;core ([#435](https://github.com/reearth/reearth/pull/435)) [`595c66`](https://github.com/reearth/reearth/commit/595c66)
- Cluster features on reearth&#x2F;core ([#430](https://github.com/reearth/reearth/pull/430)) [`92dd47`](https://github.com/reearth/reearth/commit/92dd47)
- 3D Tiles infobox on reearth&#x2F;core ([#433](https://github.com/reearth/reearth/pull/433)) [`b4afd7`](https://github.com/reearth/reearth/commit/b4afd7)
- GeoJSON with resource appearance on reearth&#x2F;core ([#432](https://github.com/reearth/reearth/pull/432)) [`464d67`](https://github.com/reearth/reearth/commit/464d67)
- Dnd layer on reearth&#x2F;core ([#424](https://github.com/reearth/reearth/pull/424)) [`75e6a7`](https://github.com/reearth/reearth/commit/75e6a7)
- WAS bug on reearth&#x2F;core ([#416](https://github.com/reearth/reearth/pull/416)) [`045274`](https://github.com/reearth/reearth/commit/045274)
- Cannot select features on reearth&#x2F;core ([#414](https://github.com/reearth/reearth/pull/414)) [`f1a8dd`](https://github.com/reearth/reearth/commit/f1a8dd)
- Support csv value string on reearth&#x2F;core ([#415](https://github.com/reearth/reearth/pull/415)) [`4033f7`](https://github.com/reearth/reearth/commit/4033f7)
- MoveTo widget to empty page on mobile on reearth&#x2F;core ([#413](https://github.com/reearth/reearth/pull/413)) [`07a935`](https://github.com/reearth/reearth/commit/07a935)
- Widget align system on mobile ([#409](https://github.com/reearth/reearth/pull/409)) [`fce1ad`](https://github.com/reearth/reearth/commit/fce1ad)
- Increment runTimes on preinit ([#410](https://github.com/reearth/reearth/pull/410)) [`797020`](https://github.com/reearth/reearth/commit/797020)
- Plugin instance runTimes property ([#405](https://github.com/reearth/reearth/pull/405)) [`a06434`](https://github.com/reearth/reearth/commit/a06434)
- Make iframe width or height assignable without the other ([#403](https://github.com/reearth/reearth/pull/403)) [`e8647a`](https://github.com/reearth/reearth/commit/e8647a)
- Cache ComputedFeature on reearth&#x2F;core ([#396](https://github.com/reearth/reearth/pull/396)) [`775a8a`](https://github.com/reearth/reearth/commit/775a8a)
- Missing type of API modal &amp; popup options.  ([#400](https://github.com/reearth/reearth/pull/400)) [`a68b24`](https://github.com/reearth/reearth/commit/a68b24)
- Visualizer on reearth&#x2F;core ([#395](https://github.com/reearth/reearth/pull/395)) [`7ba0db`](https://github.com/reearth/reearth/commit/7ba0db)
- Update mvt dynamically when appearance is updated ([#393](https://github.com/reearth/reearth/pull/393)) [`7ca5d1`](https://github.com/reearth/reearth/commit/7ca5d1)

#### 📖 Documentation

- Add diagrams of core architecture ([#396](https://github.com/reearth/reearth/pull/396)) [`2dfa6b`](https://github.com/reearth/reearth/commit/2dfa6b)

#### ⚡️ Performance

- Consider geojson as both delegate and active data type in reearth&#x2F;core ([#608](https://github.com/reearth/reearth/pull/608)) [`94bc59`](https://github.com/reearth/reearth/commit/94bc59)
- Improve styling in MVT on reearth&#x2F;core ([#574](https://github.com/reearth/reearth/pull/574)) [`8ced77`](https://github.com/reearth/reearth/commit/8ced77)
- Improve 3dtiles performance and disable requestRenderMode on reearth&#x2F;core ([#568](https://github.com/reearth/reearth/pull/568)) [`e645ec`](https://github.com/reearth/reearth/commit/e645ec)
- Reduce style evaluator memory signature in reearth&#x2F;core ([#563](https://github.com/reearth/reearth/pull/563)) [`f74b56`](https://github.com/reearth/reearth/commit/f74b56)
- Use private modifier on evaluator on reearth&#x2F;core ([#543](https://github.com/reearth/reearth/pull/543)) [`532c66`](https://github.com/reearth/reearth/commit/532c66)
- Improve regexp on reearth&#x2F;core ([#533](https://github.com/reearth/reearth/pull/533)) [`ca7b05`](https://github.com/reearth/reearth/commit/ca7b05)
- Stop synchronizing features in MVT on reearth&#x2F;core ([#521](https://github.com/reearth/reearth/pull/521)) [`82ae2c`](https://github.com/reearth/reearth/commit/82ae2c)
- Copy lazy layer lazily ([#517](https://github.com/reearth/reearth/pull/517)) [`4a5ba4`](https://github.com/reearth/reearth/commit/4a5ba4)
- Improve mvt rendering on reearth&#x2F;core ([#501](https://github.com/reearth/reearth/pull/501)) [`8a681d`](https://github.com/reearth/reearth/commit/8a681d)
- Compute features concurrently on reearth&#x2F;core ([#499](https://github.com/reearth/reearth/pull/499)) [`6448ba`](https://github.com/reearth/reearth/commit/6448ba)
- Improve skipping computing process for 3dtiles on reearth&#x2F;core ([#491](https://github.com/reearth/reearth/pull/491)) [`253b58`](https://github.com/reearth/reearth/commit/253b58)
- Improve 3dtiles features calculation on reearth&#x2F;core ([#489](https://github.com/reearth/reearth/pull/489)) [`1204a6`](https://github.com/reearth/reearth/commit/1204a6)
- Improve expression cache strategy on reearth&#x2F;core ([#488](https://github.com/reearth/reearth/pull/488)) [`324e28`](https://github.com/reearth/reearth/commit/324e28)
- Cache AST for evaluator on reearth&#x2F;core ([#473](https://github.com/reearth/reearth/pull/473)) [`da6bb3`](https://github.com/reearth/reearth/commit/da6bb3)
- Improve blink when feature is updated on reearth&#x2F;core ([#429](https://github.com/reearth/reearth/pull/429)) [`c10a67`](https://github.com/reearth/reearth/commit/c10a67)

#### ✨ Refactor

- Replace &quot;team&quot; word  related to Team state with &quot;workspace&quot; ([#607](https://github.com/reearth/reearth/pull/607)) [`fb254b`](https://github.com/reearth/reearth/commit/fb254b)

#### Miscellaneous Tasks

- Remove redundant workflows [`f2685d`](https://github.com/reearth/reearth/commit/f2685d)
- Remove inter-dependency of web and server workflows [`999e73`](https://github.com/reearth/reearth/commit/999e73)
- Change codeowner [`bffa05`](https://github.com/reearth/reearth/commit/bffa05)
- Upgrade eslint [`d55795`](https://github.com/reearth/reearth/commit/d55795)
- Update dependency cesium to v1.104.0 ([#594](https://github.com/reearth/reearth/pull/594)) [`c57839`](https://github.com/reearth/reearth/commit/c57839)
- Fix storybook is not working ([#536](https://github.com/reearth/reearth/pull/536) [`d52124`](https://github.com/reearth/reearth/commit/d52124)
- Rename asset dir to avoid conflicts with backend API endpoints [`1d9455`](https://github.com/reearth/reearth/commit/1d9455)
- Add offline_access auth scope to support refresh tokens with built-in auth server ([#425](https://github.com/reearth/reearth/pull/425) [`2a2af1`](https://github.com/reearth/reearth/commit/2a2af1)
- Upgrade dependencies ([#391](https://github.com/reearth/reearth/pull/391)) [`7280af`](https://github.com/reearth/reearth/commit/7280af)

#### 

- Add updated SECURITY.md [`bc31ce`](https://github.com/reearth/reearth/commit/bc31ce)
- Update to 0.15.0 for the release [`d9b693`](https://github.com/reearth/reearth/commit/d9b693)
- Mono repo — moving reearth-web [`3b1d8d`](https://github.com/reearth/reearth/commit/3b1d8d)
- Chore(server(: upgrade golangci-lint to v1.51 [`9c8714`](https://github.com/reearth/reearth/commit/9c8714)

### ci

#### 🔧 Bug Fixes

- Set deafault working directory in build-docker-image [`c57909`](https://github.com/reearth/reearth/commit/c57909)
- Path director fix on docker build and refactor on deply web nightly [`7c72ea`](https://github.com/reearth/reearth/commit/7c72ea)

#### Miscellaneous Tasks

- Add conditional in ci-web and ci-server for renovate commits ([#406](https://github.com/reearth/reearth/pull/406)) [`cc024b`](https://github.com/reearth/reearth/commit/cc024b)
- Invoke workflows with their name [`fbbb76`](https://github.com/reearth/reearth/commit/fbbb76)

### core

#### 🚀 Features

- Support shapefile in reearth&#x2F;core ([#420](https://github.com/reearth/reearth/pull/420)) [`508502`](https://github.com/reearth/reearth/commit/508502)
- Support gpx in reearth&#x2F;core ([#423](https://github.com/reearth/reearth/pull/423)) [`b4b04b`](https://github.com/reearth/reearth/commit/b4b04b)

### 

#### 🚀 Features

- Support polyline and point on reearth&#x2F;core ([#606](https://github.com/reearth/reearth/pull/606)) [`abd9a3`](https://github.com/reearth/reearth/commit/abd9a3)
- Add remembering last open workspace functionality ([#598](https://github.com/reearth/reearth/pull/598)) [`968a9d`](https://github.com/reearth/reearth/commit/968a9d)
- Support TileMapService(TMS) on reearth&#x2F;core ([#604](https://github.com/reearth/reearth/pull/604)) [`284b35`](https://github.com/reearth/reearth/commit/284b35)
- Support styling color &amp; show for 3dtiles model ([#599](https://github.com/reearth/reearth/pull/599)) [`a82ca2`](https://github.com/reearth/reearth/commit/a82ca2)
- Support tiles data type on reearth&#x2F;core ([#597](https://github.com/reearth/reearth/pull/597) [`ae5c49`](https://github.com/reearth/reearth/commit/ae5c49)
- Support classificationType property in some feature on reearth&#x2F;core ([#593](https://github.com/reearth/reearth/pull/593)) [`897868`](https://github.com/reearth/reearth/commit/897868)
- Upgrade cesium-mvt-imagery-provider ([#591](https://github.com/reearth/reearth/pull/591)) [`d01d01`](https://github.com/reearth/reearth/commit/d01d01)
- Upgrade cesium-mvt-imagery-provider ([#588](https://github.com/reearth/reearth/pull/588)) [`e138bd`](https://github.com/reearth/reearth/commit/e138bd)
- Add more styling properties to resource appearance in reearth&#x2F;core ([#586](https://github.com/reearth/reearth/pull/586)) [`8f3625`](https://github.com/reearth/reearth/commit/8f3625)
- Add built field in scene in plugin api ([#584](https://github.com/reearth/reearth/pull/584)) [`e8050c`](https://github.com/reearth/reearth/commit/e8050c)
- Add scene light ([#576](https://github.com/reearth/reearth/pull/576)) [`43c2b8`](https://github.com/reearth/reearth/commit/43c2b8)
- Upgrade mvt lib ([#575](https://github.com/reearth/reearth/pull/575)) [`584cee`](https://github.com/reearth/reearth/commit/584cee)
- Option to unselect layer when click infobox close ([#564](https://github.com/reearth/reearth/pull/564)) [`f2b2f2`](https://github.com/reearth/reearth/commit/f2b2f2)
- Add alpha property to raster appearance on reearth&#x2F;core ([#555](https://github.com/reearth/reearth/pull/555)) [`541f30`](https://github.com/reearth/reearth/commit/541f30)
- Support styling for point cloud on reearth&#x2F;core ([#549](https://github.com/reearth/reearth/pull/549)) [`f410d6`](https://github.com/reearth/reearth/commit/f410d6)
- Add defaultContent property for infobox in plugin API on reearth&#x2F;core ([#538](https://github.com/reearth/reearth/pull/538)) [`31ba31`](https://github.com/reearth/reearth/commit/31ba31)
- Add updateClockOnLoad  to data on reearth&#x2F;core ([#539](https://github.com/reearth/reearth/pull/539)) [`3653e2`](https://github.com/reearth/reearth/commit/3653e2)
- Support gltf data type on reearth&#x2F;core ([#535](https://github.com/reearth/reearth/pull/535)) [`e086be`](https://github.com/reearth/reearth/commit/e086be)
- Support ga4 ([#509](https://github.com/reearth/reearth/pull/509)) [`39bfb0`](https://github.com/reearth/reearth/commit/39bfb0)
- Select MVT feature on reearth&#x2F;core ([#527](https://github.com/reearth/reearth/pull/527)) [`de605d`](https://github.com/reearth/reearth/commit/de605d)
- Add parameters property to data on reearth&#x2F;core ([#520](https://github.com/reearth/reearth/pull/520)) [`c698eb`](https://github.com/reearth/reearth/commit/c698eb)
- Layer select event in reearth&#x2F;core ([#470](https://github.com/reearth/reearth/pull/470)) [`fb22e6`](https://github.com/reearth/reearth/commit/fb22e6)
- Show field in appearances of reearth&#x2F;core ([#469](https://github.com/reearth/reearth/pull/469)) [`819eb6`](https://github.com/reearth/reearth/commit/819eb6)
- ImageSizeInMeters field in marker proeperty ([#511](https://github.com/reearth/reearth/pull/511)) [`290cb7`](https://github.com/reearth/reearth/commit/290cb7)
- Override clock from scene setting on reearth&#x2F;core ([#505](https://github.com/reearth/reearth/pull/505)) [`01bffd`](https://github.com/reearth/reearth/commit/01bffd)
- Support features for CZML on reearth&#x2F;core ([#506](https://github.com/reearth/reearth/pull/506)) [`e5c160`](https://github.com/reearth/reearth/commit/e5c160)
- Update cesium ([#503](https://github.com/reearth/reearth/pull/503)) [`5a649f`](https://github.com/reearth/reearth/commit/5a649f)
- Color blend mode in tileset on reearth&#x2F;core ([#496](https://github.com/reearth/reearth/pull/496)) [`ca43dc`](https://github.com/reearth/reearth/commit/ca43dc)
- Change brand images and colors at the root page ([#495](https://github.com/reearth/reearth/pull/495))Co-authored-by: rot1024 &lt;aayhrot@gmail.com&gt; [`4f07b9`](https://github.com/reearth/reearth/commit/4f07b9)
- Add builtin clipping box on reearth&#x2F;core ([#487](https://github.com/reearth/reearth/pull/487)) [`63bd4f`](https://github.com/reearth/reearth/commit/63bd4f)
- ExtrudedHeight for polygon on reearth&#x2F;core ([#486](https://github.com/reearth/reearth/pull/486)) [`523d35`](https://github.com/reearth/reearth/commit/523d35)
- Support resource entity layerId on reearth&#x2F;core ([#485](https://github.com/reearth/reearth/pull/485)) [`7bd7c5`](https://github.com/reearth/reearth/commit/7bd7c5)
- Support distanceDisplayCondition on reearth&#x2F;core ([#475](https://github.com/reearth/reearth/pull/475)) [`ce8270`](https://github.com/reearth/reearth/commit/ce8270)
- Support GeoRSS and gml in reearth&#x2F;core ([#455](https://github.com/reearth/reearth/pull/455)) [`58c25b`](https://github.com/reearth/reearth/commit/58c25b)
- Add htmlblock on reearth&#x2F;core ([#454](https://github.com/reearth/reearth/pull/454)) [`1b37e0`](https://github.com/reearth/reearth/commit/1b37e0)
- Support clipping box direction on reearth&#x2F;core ([#467](https://github.com/reearth/reearth/pull/467)) [`70f74e`](https://github.com/reearth/reearth/commit/70f74e)
- Add sampleTerrainHeight on reearht&#x2F;core ([#466](https://github.com/reearth/reearth/pull/466)) [`55334e`](https://github.com/reearth/reearth/commit/55334e)
- Get brand from config ([#457](https://github.com/reearth/reearth/pull/457)) [`d35361`](https://github.com/reearth/reearth/commit/d35361)
- Support timeline on mobile on reearth&#x2F;core ([#462](https://github.com/reearth/reearth/pull/462)) [`efeaf4`](https://github.com/reearth/reearth/commit/efeaf4)
- Use visible field on reearth&#x2F;core ([#456](https://github.com/reearth/reearth/pull/456)) [`333610`](https://github.com/reearth/reearth/commit/333610)
- Add htmlblock to built-in plugin ([#384](https://github.com/reearth/reearth/pull/384)) [`51c79a`](https://github.com/reearth/reearth/commit/51c79a)
- Add override, replace and delete plugin API on reearth&#x2F;core ([#451](https://github.com/reearth/reearth/pull/451)) [`2e1c41`](https://github.com/reearth/reearth/commit/2e1c41)
- Selecting imagery features ([#450](https://github.com/reearth/reearth/pull/450)) [`f24ef5`](https://github.com/reearth/reearth/commit/f24ef5)
- Support interval fetching data on reearth&#x2F;core ([#449](https://github.com/reearth/reearth/pull/449)) [`406174`](https://github.com/reearth/reearth/commit/406174)
- Support select feature on reearth&#x2F;core ([#445](https://github.com/reearth/reearth/pull/445)) [`3174b1`](https://github.com/reearth/reearth/commit/3174b1)
- Use experimental core flag ([#448](https://github.com/reearth/reearth/pull/448)) [`b04294`](https://github.com/reearth/reearth/commit/b04294)
- Support time series features on reearth&#x2F;core ([#446](https://github.com/reearth/reearth/pull/446)) [`8fc9b6`](https://github.com/reearth/reearth/commit/8fc9b6)
- Scene property to enable VR mode ([#444](https://github.com/reearth/reearth/pull/444)) [`3d35aa`](https://github.com/reearth/reearth/commit/3d35aa)
- Support general transit feed ([#408](https://github.com/reearth/reearth/pull/408)) [`49b4a4`](https://github.com/reearth/reearth/commit/49b4a4)
- Support osm data type on reearth&#x2F;core ([#431](https://github.com/reearth/reearth/pull/431)) [`0d4e0b`](https://github.com/reearth/reearth/commit/0d4e0b)
- Label background color and padding property for markers ([#426](https://github.com/reearth/reearth/pull/426) [`72cd0d`](https://github.com/reearth/reearth/commit/72cd0d)
- Support entity base flyTo on reearth&#x2F;core ([#419](https://github.com/reearth/reearth/pull/419)) [`3060cf`](https://github.com/reearth/reearth/commit/3060cf)
- Support overriding czml appearance on reearth&#x2F;core ([#421](https://github.com/reearth/reearth/pull/421)) [`e62f4d`](https://github.com/reearth/reearth/commit/e62f4d)
- Support kml on reearth&#x2F;core ([#422](https://github.com/reearth/reearth/pull/422)) [`052daf`](https://github.com/reearth/reearth/commit/052daf)
- Support json properties on reearth&#x2F;core ([#412](https://github.com/reearth/reearth/pull/412)) [`ac7986`](https://github.com/reearth/reearth/commit/ac7986)
- Connect reearth&#x2F;core with existence pages ([#401](https://github.com/reearth/reearth/pull/401)) [`0735c0`](https://github.com/reearth/reearth/commit/0735c0)
- Add runTimes property to PluginInstance ([#404](https://github.com/reearth/reearth/pull/404)) [`17d787`](https://github.com/reearth/reearth/commit/17d787)
- Support plugin system on reearth&#x2F;core ([#399](https://github.com/reearth/reearth/pull/399)) [`bab9e6`](https://github.com/reearth/reearth/commit/bab9e6)
- Add selectedFeature and selectedComputedFeature on reearth&#x2F;core ([#398](https://github.com/reearth/reearth/pull/398)) [`474b34`](https://github.com/reearth/reearth/commit/474b34)
- Set modal above popup ([#397](https://github.com/reearth/reearth/pull/397)) [`ff47c5`](https://github.com/reearth/reearth/commit/ff47c5)
- Port 2d navigator to reearth&#x2F;core ([#394](https://github.com/reearth/reearth/pull/394)) [`07a6b4`](https://github.com/reearth/reearth/commit/07a6b4)
- Core&#x2F;Visualizer without plugins ([#372](https://github.com/reearth/reearth/pull/372)) [`f97c38`](https://github.com/reearth/reearth/commit/f97c38)
- Reearth style language ([#384](https://github.com/reearth/reearth/pull/384)) [`a828ac`](https://github.com/reearth/reearth/commit/a828ac)
- Support 3dtiles on reearth&#x2F;core ([#392](https://github.com/reearth/reearth/pull/392)) [`e8068f`](https://github.com/reearth/reearth/commit/e8068f)
- Support MVT on reearth&#x2F;core ([#388](https://github.com/reearth/reearth/pull/388)) [`cac89c`](https://github.com/reearth/reearth/commit/cac89c)
- Support WMS on reearth&#x2F;core ([#387](https://github.com/reearth/reearth/pull/387)) [`666c1b`](https://github.com/reearth/reearth/commit/666c1b)
- Support CZML in reearth&#x2F;core ([#383](https://github.com/reearth/reearth/pull/383)) [`f44d98`](https://github.com/reearth/reearth/commit/f44d98)
- Plugin api client storage ([#376](https://github.com/reearth/reearth/pull/376)) [`4f36ad`](https://github.com/reearth/reearth/commit/4f36ad)
- Support csv on the reearth&#x2F;core ([#382](https://github.com/reearth/reearth/pull/382)) [`a8f5bf`](https://github.com/reearth/reearth/commit/a8f5bf)

#### 🔧 Bug Fixes

- Merge conflict resolved [`39a37a`](https://github.com/reearth/reearth/commit/39a37a)
- Use core hook may get value incorrectly ([#609](https://github.com/reearth/reearth/pull/609)) [`d87e90`](https://github.com/reearth/reearth/commit/d87e90)
- Resolve endsWith is not a function error with style in reearth&#x2F;core ([#603](https://github.com/reearth/reearth/pull/603)) [`15fec3`](https://github.com/reearth/reearth/commit/15fec3)
- Add % as special case for &#x60;Number()&#x60; in style lang of reearth&#x2F;core ([#601](https://github.com/reearth/reearth/pull/601)) [`5a4a3d`](https://github.com/reearth/reearth/commit/5a4a3d)
- Rename moveWidget to onMoveWidget on reearth&#x2F;core ([#600](https://github.com/reearth/reearth/pull/600)) [`6c06ab`](https://github.com/reearth/reearth/commit/6c06ab)
- Enable splash screen in preview page ([#596](https://github.com/reearth/reearth/pull/596) [`e1f5ac`](https://github.com/reearth/reearth/commit/e1f5ac)
- Timeline scroll should be fixed in initial render ([#571](https://github.com/reearth/reearth/pull/571)) [`6a0aed`](https://github.com/reearth/reearth/commit/6a0aed)
- Remove default height reference for modelGraphics in reearth&#x2F;core ([#592](https://github.com/reearth/reearth/pull/592)) [`961b46`](https://github.com/reearth/reearth/commit/961b46)
- Remove duplication of feature entity in reearth&#x2F;core ([#590](https://github.com/reearth/reearth/pull/590)) [`8cc03e`](https://github.com/reearth/reearth/commit/8cc03e)
- Prevent selecting not shown feature for MVT on reearth&#x2F;core ([#589](https://github.com/reearth/reearth/pull/589)) [`41816d`](https://github.com/reearth/reearth/commit/41816d)
- Infobox for 3dtiles feature ([#587](https://github.com/reearth/reearth/pull/587)) [`70cfdb`](https://github.com/reearth/reearth/commit/70cfdb)
- Use layer id with feature id for entity id on reearth&#x2F;core ([#585](https://github.com/reearth/reearth/pull/585)) [`24cc88`](https://github.com/reearth/reearth/commit/24cc88)
- Can&[#39](https://github.com/reearth/reearth/pull/39);t set iframe&[#39](https://github.com/reearth/reearth/pull/39);s width or height individually in crust  ([#583](https://github.com/reearth/reearth/pull/583)) [`3494b4`](https://github.com/reearth/reearth/commit/3494b4)
- MVT isn&[#39](https://github.com/reearth/reearth/pull/39);t render correctly on reearth&#x2F;core ([#582](https://github.com/reearth/reearth/pull/582)) [`988d1a`](https://github.com/reearth/reearth/commit/988d1a)
- Sample terrain height API cannot properly return promise ([#581](https://github.com/reearth/reearth/pull/581)) [`72eafd`](https://github.com/reearth/reearth/commit/72eafd)
- Change logical operator evaluation behaviour in reearth&#x2F;core ([#580](https://github.com/reearth/reearth/pull/580)) [`3d23ce`](https://github.com/reearth/reearth/commit/3d23ce)
- Infobox style collapse for long names ([#578](https://github.com/reearth/reearth/pull/578)) [`04775d`](https://github.com/reearth/reearth/commit/04775d)
- Override currentTime on timeline widget when time is updated by CZML on reearth&#x2F;core ([#579](https://github.com/reearth/reearth/pull/579)) [`dbe2e1`](https://github.com/reearth/reearth/commit/dbe2e1)
- Rename scene light properties [`8b3a18`](https://github.com/reearth/reearth/commit/8b3a18)
- Add none cesium value for heightReference ([#577](https://github.com/reearth/reearth/pull/577)) [`2c92e3`](https://github.com/reearth/reearth/commit/2c92e3)
- JSON Path condition on reearth&#x2F;core ([#572](https://github.com/reearth/reearth/pull/572)) [`5f7024`](https://github.com/reearth/reearth/commit/5f7024)
- Update html head [`f1780f`](https://github.com/reearth/reearth/commit/f1780f)
- Sync selected feature with layer api on reearth&#x2F;core ([#570](https://github.com/reearth/reearth/pull/570)) [`ded6b0`](https://github.com/reearth/reearth/commit/ded6b0)
- Imagery index on reearth&#x2F;core ([#569](https://github.com/reearth/reearth/pull/569)) [`6b233f`](https://github.com/reearth/reearth/commit/6b233f)
- Revert add height to polygon in reearth&#x2F;core ([#566](https://github.com/reearth/reearth/pull/566)) [`1f2f74`](https://github.com/reearth/reearth/commit/1f2f74)
- Add clamp as default height reference for model in reearth&#x2F;core ([#567](https://github.com/reearth/reearth/pull/567)) [`517386`](https://github.com/reearth/reearth/commit/517386)
- Minimum timeline range on reearth&#x2F;core ([#565](https://github.com/reearth/reearth/pull/565)) [`975c79`](https://github.com/reearth/reearth/commit/975c79)
- Infobox html block styling ([#562](https://github.com/reearth/reearth/pull/562)) [`32b248`](https://github.com/reearth/reearth/commit/32b248)
- Imagery layer tile index on reearth&#x2F;core ([#561](https://github.com/reearth/reearth/pull/561) [`25bdff`](https://github.com/reearth/reearth/commit/25bdff)
- Add &quot;disabled&quot; as shadow mode on reearth&#x2F;core ([#560](https://github.com/reearth/reearth/pull/560)) [`1632ea`](https://github.com/reearth/reearth/commit/1632ea)
- Support multi layers in MVT on reearth&#x2F;core ([#559](https://github.com/reearth/reearth/pull/559)) [`9c8c88`](https://github.com/reearth/reearth/commit/9c8c88)
- Html block margin and height [`09f6ef`](https://github.com/reearth/reearth/commit/09f6ef)
- UpdateClockOnLoad condition on reearth&#x2F;core ([#558](https://github.com/reearth/reearth/pull/558)) [`f9076a`](https://github.com/reearth/reearth/commit/f9076a)
- Add height to polygon in reearth&#x2F;core ([#557](https://github.com/reearth/reearth/pull/557)) [`a84dbe`](https://github.com/reearth/reearth/commit/a84dbe)
- Use computed feature on resource on reearth&#x2F;core ([#556](https://github.com/reearth/reearth/pull/556)) [`006f40`](https://github.com/reearth/reearth/commit/006f40)
- Point cloud dirty check on reearth&#x2F;core ([#554](https://github.com/reearth/reearth/pull/554)) [`f95d27`](https://github.com/reearth/reearth/commit/f95d27)
- Support default infobox and selection indicator on imagery layer on reearth&#x2F;core ([#553](https://github.com/reearth/reearth/pull/553)) [`1946b3`](https://github.com/reearth/reearth/commit/1946b3)
- Scrollbar in timeline widget always showing ([#550](https://github.com/reearth/reearth/pull/550)) [`84d63d`](https://github.com/reearth/reearth/commit/84d63d)
- Clipping box on point cloud on reearth&#x2F;core ([#552](https://github.com/reearth/reearth/pull/552)) [`fb4455`](https://github.com/reearth/reearth/commit/fb4455)
- Htmlblock on safari ([#548](https://github.com/reearth/reearth/pull/548)) [`392fc7`](https://github.com/reearth/reearth/commit/392fc7)
- Use attributes as default content for infobox on reearth&#x2F;core ([#547](https://github.com/reearth/reearth/pull/547)) [`be3718`](https://github.com/reearth/reearth/commit/be3718)
- Deleting feature process on reearth&#x2F;core ([#546](https://github.com/reearth/reearth/pull/546)) [`39ecaf`](https://github.com/reearth/reearth/commit/39ecaf)
- Trigger select event when featureId is changed on reearth&#x2F;core ([#545](https://github.com/reearth/reearth/pull/545)) [`6f5401`](https://github.com/reearth/reearth/commit/6f5401)
- Draw polylines on polygon on reearth&#x2F;core ([#544](https://github.com/reearth/reearth/pull/544)) [`d35329`](https://github.com/reearth/reearth/commit/d35329)
- Infobox html color ([#534](https://github.com/reearth/reearth/pull/534)) [`c0e0a6`](https://github.com/reearth/reearth/commit/c0e0a6)
- Use overriddenLayers to get infobox on reearth&#x2F;core ([#541](https://github.com/reearth/reearth/pull/541)) [`c4e9db`](https://github.com/reearth/reearth/commit/c4e9db)
- Fly to multiple entities added by a layer on reearth&#x2F;core ([#540](https://github.com/reearth/reearth/pull/540)) [`95b3d3`](https://github.com/reearth/reearth/commit/95b3d3)
- Add polyfill for requestIdleCallback in reearth&#x2F;core ([#537](https://github.com/reearth/reearth/pull/537)) [`c4722f`](https://github.com/reearth/reearth/commit/c4722f)
- Overriding timeline behavior on reearth&#x2F;core ([#532](https://github.com/reearth/reearth/pull/532)) [`890dae`](https://github.com/reearth/reearth/commit/890dae)
- Timeline bug on reearth&#x2F;core ([#531](https://github.com/reearth/reearth/pull/531)) [`572678`](https://github.com/reearth/reearth/commit/572678)
- Some error on reearth&#x2F;core ([#530](https://github.com/reearth/reearth/pull/530)) [`4ee0b6`](https://github.com/reearth/reearth/commit/4ee0b6)
- Abort fetching on data atom on reearth&#x2F;core ([#529](https://github.com/reearth/reearth/pull/529)) [`e82c88`](https://github.com/reearth/reearth/commit/e82c88)
- Timeline behavior on reearth&#x2F;core ([#528](https://github.com/reearth/reearth/pull/528)) [`461d06`](https://github.com/reearth/reearth/commit/461d06)
- Parse csv numeric strings as numbers in reearth&#x2F;core ([#526](https://github.com/reearth/reearth/pull/526)) [`9c890f`](https://github.com/reearth/reearth/commit/9c890f)
- Parse hyphen as reserved word in property key on reearth&#x2F;core ([#525](https://github.com/reearth/reearth/pull/525)) [`d49058`](https://github.com/reearth/reearth/commit/d49058)
- Condition for CZML on reearth&#x2F;core ([#524](https://github.com/reearth/reearth/pull/524)) [`3c61ca`](https://github.com/reearth/reearth/commit/3c61ca)
- Czml style for marker on reearth&#x2F;core ([#523](https://github.com/reearth/reearth/pull/523)) [`67dccb`](https://github.com/reearth/reearth/commit/67dccb)
- Use default block for entity on reearth&#x2F;core ([#522](https://github.com/reearth/reearth/pull/522)) [`ecd09a`](https://github.com/reearth/reearth/commit/ecd09a)
- Remove copyLazyLayer on reearth&#x2F;core ([#519](https://github.com/reearth/reearth/pull/519)) [`e17218`](https://github.com/reearth/reearth/commit/e17218)
- Copying lazy layers undefined behavior on reearth&#x2F;core ([#516](https://github.com/reearth/reearth/pull/516)) [`7b2c5e`](https://github.com/reearth/reearth/commit/7b2c5e)
- Copy lazy layer on plugin on reearth&#x2F;core ([#515](https://github.com/reearth/reearth/pull/515)) [`046a3d`](https://github.com/reearth/reearth/commit/046a3d)
- Attach style dynamically in resource on reearth&#x2F;core ([#514](https://github.com/reearth/reearth/pull/514)) [`81d291`](https://github.com/reearth/reearth/commit/81d291)
- Infinite infobox in CZML on reearth&#x2F;core ([#513](https://github.com/reearth/reearth/pull/513)) [`3e49a1`](https://github.com/reearth/reearth/commit/3e49a1)
- Selecting resource feature behavior on reearth&#x2F;core ([#512](https://github.com/reearth/reearth/pull/512)) [`35f2c2`](https://github.com/reearth/reearth/commit/35f2c2)
- Prevent unnecessary render on timeline on reearth&#x2F;core ([#510](https://github.com/reearth/reearth/pull/510)) [`7cb9d5`](https://github.com/reearth/reearth/commit/7cb9d5)
- Parse reserved word when property name includes reserved word on reearth&#x2F;core ([#508](https://github.com/reearth/reearth/pull/508)) [`aa247e`](https://github.com/reearth/reearth/commit/aa247e)
- Pass engine meta on plugin editor ([#507](https://github.com/reearth/reearth/pull/507)) [`9cd1b5`](https://github.com/reearth/reearth/commit/9cd1b5)
- Mvt line width on reearth&#x2F;core ([#504](https://github.com/reearth/reearth/pull/504)) [`c1a939`](https://github.com/reearth/reearth/commit/c1a939)
- It should not render entity when coordinate is undefined on reearth&#x2F;core [`23b6c6`](https://github.com/reearth/reearth/commit/23b6c6)
- Allow enter ground option for clipping box on reearth&#x2F;core ([#500](https://github.com/reearth/reearth/pull/500)) [`f8e129`](https://github.com/reearth/reearth/commit/f8e129)
- Clip area with clipping box on reearth&#x2F;core ([#498](https://github.com/reearth/reearth/pull/498)) [`4f647f`](https://github.com/reearth/reearth/commit/4f647f)
- Mvt cache on reearth&#x2F;core ([#497](https://github.com/reearth/reearth/pull/497)) [`ba2c7e`](https://github.com/reearth/reearth/commit/ba2c7e)
- Recreate no feature component when data url is changed on reearth&#x2F;core ([#494](https://github.com/reearth/reearth/pull/494)) [`93c805`](https://github.com/reearth/reearth/commit/93c805)
- Ignore cesium ion token when it is empty [`6648be`](https://github.com/reearth/reearth/commit/6648be)
- Feature type fix gpx on reearth&#x2F;core ([#493](https://github.com/reearth/reearth/pull/493)) [`51b8e8`](https://github.com/reearth/reearth/commit/51b8e8)
- Infobox error after layer delete ([#492](https://github.com/reearth/reearth/pull/492)) [`8a59dd`](https://github.com/reearth/reearth/commit/8a59dd)
- Delete bug that deletes necessary layer on reearth&#x2F;core ([#484](https://github.com/reearth/reearth/pull/484)) [`a0f48f`](https://github.com/reearth/reearth/commit/a0f48f)
- Force update when some data properties are updated on reearth&#x2F;core ([#483](https://github.com/reearth/reearth/pull/483)) [`867238`](https://github.com/reearth/reearth/commit/867238)
- Revert appearance to initial value when appearance is undefined on reearth&#x2F;core ([#482](https://github.com/reearth/reearth/pull/482)) [`76b6f4`](https://github.com/reearth/reearth/commit/76b6f4)
- Entity id is duplicated error on reearth&#x2F;core ([#481](https://github.com/reearth/reearth/pull/481)) [`a2b0b6`](https://github.com/reearth/reearth/commit/a2b0b6)
- Infobox property is undefined error on reearth&#x2F;core ([#480](https://github.com/reearth/reearth/pull/480)) [`2b8b07`](https://github.com/reearth/reearth/commit/2b8b07)
- Layers override behavior on rearth&#x2F;core ([#479](https://github.com/reearth/reearth/pull/479)) [`faff37`](https://github.com/reearth/reearth/commit/faff37)
- Error handling for time interval on reearth&#x2F;core ([#478](https://github.com/reearth/reearth/pull/478)) [`edf546`](https://github.com/reearth/reearth/commit/edf546)
- Overridden layers api on reearth&#x2F;core ([#477](https://github.com/reearth/reearth/pull/477)) [`abaade`](https://github.com/reearth/reearth/commit/abaade)
- Add properties in vehicle point for gtfs ([#476](https://github.com/reearth/reearth/pull/476)) [`1058ab`](https://github.com/reearth/reearth/commit/1058ab)
- Color function on reearth&#x2F;core ([#474](https://github.com/reearth/reearth/pull/474)) [`7d9fca`](https://github.com/reearth/reearth/commit/7d9fca)
- Errors when many functions are created in plugins ([#471](https://github.com/reearth/reearth/pull/471)) [`ebb50d`](https://github.com/reearth/reearth/commit/ebb50d)
- Coordinates for csv on reearth&#x2F;core ([#472](https://github.com/reearth/reearth/pull/472)) [`4a6473`](https://github.com/reearth/reearth/commit/4a6473)
- Lint and type error [`d00b9b`](https://github.com/reearth/reearth/commit/d00b9b)
- Suppress screen flicker when judging useCore [`b1852d`](https://github.com/reearth/reearth/commit/b1852d)
- Widget area has margin even if no widgets, disable widget area transition in built scene [`b51569`](https://github.com/reearth/reearth/commit/b51569)
- Cesium crashes when VR mode is false [`efa3fd`](https://github.com/reearth/reearth/commit/efa3fd)
- 3dtiles overriding appearance behavior on reearth&#x2F;core ([#468](https://github.com/reearth/reearth/pull/468)) [`8c48bc`](https://github.com/reearth/reearth/commit/8c48bc)
- Support visible and dynamic plane in clipping box on reearth&#x2F;core ([#465](https://github.com/reearth/reearth/pull/465)) [`4c89aa`](https://github.com/reearth/reearth/commit/4c89aa)
- SelectedFeature for 3dtiles on reearth&#x2F;core ([#463](https://github.com/reearth/reearth/pull/463)) [`cd1777`](https://github.com/reearth/reearth/commit/cd1777)
- Express undefined visible field on reearth&#x2F;core ([#461](https://github.com/reearth/reearth/pull/461)) [`c74630`](https://github.com/reearth/reearth/commit/c74630)
- Undefined behavior for visible field ([#460](https://github.com/reearth/reearth/pull/460)) [`c41d70`](https://github.com/reearth/reearth/commit/c41d70)
- Select entity on reearth&#x2F;core ([#458](https://github.com/reearth/reearth/pull/458)) [`bc1824`](https://github.com/reearth/reearth/commit/bc1824)
- Use default infobox on reearth&#x2F;core ([#453](https://github.com/reearth/reearth/pull/453)) [`d3fec8`](https://github.com/reearth/reearth/commit/d3fec8)
- Select event behavior on reearth&#x2F;core ([#452](https://github.com/reearth/reearth/pull/452)) [`384488`](https://github.com/reearth/reearth/commit/384488)
- Error in published page on reearth&#x2F;core ([#447](https://github.com/reearth/reearth/pull/447)) [`4c8805`](https://github.com/reearth/reearth/commit/4c8805)
- Expand timeline in initial load ([#443](https://github.com/reearth/reearth/pull/443)) [`d6a742`](https://github.com/reearth/reearth/commit/d6a742)
- Replace globe image when cesium ion token is updated ([#442](https://github.com/reearth/reearth/pull/442)) [`64ffae`](https://github.com/reearth/reearth/commit/64ffae)
- Layer fetch on reearth&#x2F;core ([#441](https://github.com/reearth/reearth/pull/441)) [`597b82`](https://github.com/reearth/reearth/commit/597b82)
- Dnd layer on reearth&#x2F;core ([#440](https://github.com/reearth/reearth/pull/440)) [`a5a2b4`](https://github.com/reearth/reearth/commit/a5a2b4)
- Disable requestRenderMode depends on widget on reearth&#x2F;core ([#439](https://github.com/reearth/reearth/pull/439)) [`12ce63`](https://github.com/reearth/reearth/commit/12ce63)
- Selected layer id is not propagated on reearth&#x2F;core ([#438](https://github.com/reearth/reearth/pull/438)) [`24993b`](https://github.com/reearth/reearth/commit/24993b)
- Handle featureId for 3dtiles and compat select plugin api on reearth&#x2F;core ([#417](https://github.com/reearth/reearth/pull/417)) [`9144ad`](https://github.com/reearth/reearth/commit/9144ad)
- Undefined behavior for resource on reearth&#x2F;core ([#437](https://github.com/reearth/reearth/pull/437)) [`3f51f2`](https://github.com/reearth/reearth/commit/3f51f2)
- Blocks cannot be displayed and updated as expected on reearth&#x2F;core ([#434](https://github.com/reearth/reearth/pull/434)) [`b5f921`](https://github.com/reearth/reearth/commit/b5f921)
- Layer appearances are not evaluated as expected ([#418](https://github.com/reearth/reearth/pull/418)) [`20382c`](https://github.com/reearth/reearth/commit/20382c)
- Support resource auto on reearth&#x2F;core ([#435](https://github.com/reearth/reearth/pull/435)) [`595c66`](https://github.com/reearth/reearth/commit/595c66)
- Cluster features on reearth&#x2F;core ([#430](https://github.com/reearth/reearth/pull/430)) [`92dd47`](https://github.com/reearth/reearth/commit/92dd47)
- 3D Tiles infobox on reearth&#x2F;core ([#433](https://github.com/reearth/reearth/pull/433)) [`b4afd7`](https://github.com/reearth/reearth/commit/b4afd7)
- GeoJSON with resource appearance on reearth&#x2F;core ([#432](https://github.com/reearth/reearth/pull/432)) [`464d67`](https://github.com/reearth/reearth/commit/464d67)
- Dnd layer on reearth&#x2F;core ([#424](https://github.com/reearth/reearth/pull/424)) [`75e6a7`](https://github.com/reearth/reearth/commit/75e6a7)
- WAS bug on reearth&#x2F;core ([#416](https://github.com/reearth/reearth/pull/416)) [`045274`](https://github.com/reearth/reearth/commit/045274)
- Cannot select features on reearth&#x2F;core ([#414](https://github.com/reearth/reearth/pull/414)) [`f1a8dd`](https://github.com/reearth/reearth/commit/f1a8dd)
- Support csv value string on reearth&#x2F;core ([#415](https://github.com/reearth/reearth/pull/415)) [`4033f7`](https://github.com/reearth/reearth/commit/4033f7)
- MoveTo widget to empty page on mobile on reearth&#x2F;core ([#413](https://github.com/reearth/reearth/pull/413)) [`07a935`](https://github.com/reearth/reearth/commit/07a935)
- Widget align system on mobile ([#409](https://github.com/reearth/reearth/pull/409)) [`fce1ad`](https://github.com/reearth/reearth/commit/fce1ad)
- Increment runTimes on preinit ([#410](https://github.com/reearth/reearth/pull/410)) [`797020`](https://github.com/reearth/reearth/commit/797020)
- Plugin instance runTimes property ([#405](https://github.com/reearth/reearth/pull/405)) [`a06434`](https://github.com/reearth/reearth/commit/a06434)
- Make iframe width or height assignable without the other ([#403](https://github.com/reearth/reearth/pull/403)) [`e8647a`](https://github.com/reearth/reearth/commit/e8647a)
- Cache ComputedFeature on reearth&#x2F;core ([#396](https://github.com/reearth/reearth/pull/396)) [`775a8a`](https://github.com/reearth/reearth/commit/775a8a)
- Missing type of API modal &amp; popup options.  ([#400](https://github.com/reearth/reearth/pull/400)) [`a68b24`](https://github.com/reearth/reearth/commit/a68b24)
- Visualizer on reearth&#x2F;core ([#395](https://github.com/reearth/reearth/pull/395)) [`7ba0db`](https://github.com/reearth/reearth/commit/7ba0db)
- Update mvt dynamically when appearance is updated ([#393](https://github.com/reearth/reearth/pull/393)) [`7ca5d1`](https://github.com/reearth/reearth/commit/7ca5d1)

#### 📖 Documentation

- Add diagrams of core architecture ([#396](https://github.com/reearth/reearth/pull/396)) [`2dfa6b`](https://github.com/reearth/reearth/commit/2dfa6b)

#### ⚡️ Performance

- Consider geojson as both delegate and active data type in reearth&#x2F;core ([#608](https://github.com/reearth/reearth/pull/608)) [`94bc59`](https://github.com/reearth/reearth/commit/94bc59)
- Improve styling in MVT on reearth&#x2F;core ([#574](https://github.com/reearth/reearth/pull/574)) [`8ced77`](https://github.com/reearth/reearth/commit/8ced77)
- Improve 3dtiles performance and disable requestRenderMode on reearth&#x2F;core ([#568](https://github.com/reearth/reearth/pull/568)) [`e645ec`](https://github.com/reearth/reearth/commit/e645ec)
- Reduce style evaluator memory signature in reearth&#x2F;core ([#563](https://github.com/reearth/reearth/pull/563)) [`f74b56`](https://github.com/reearth/reearth/commit/f74b56)
- Use private modifier on evaluator on reearth&#x2F;core ([#543](https://github.com/reearth/reearth/pull/543)) [`532c66`](https://github.com/reearth/reearth/commit/532c66)
- Improve regexp on reearth&#x2F;core ([#533](https://github.com/reearth/reearth/pull/533)) [`ca7b05`](https://github.com/reearth/reearth/commit/ca7b05)
- Stop synchronizing features in MVT on reearth&#x2F;core ([#521](https://github.com/reearth/reearth/pull/521)) [`82ae2c`](https://github.com/reearth/reearth/commit/82ae2c)
- Copy lazy layer lazily ([#517](https://github.com/reearth/reearth/pull/517)) [`4a5ba4`](https://github.com/reearth/reearth/commit/4a5ba4)
- Improve mvt rendering on reearth&#x2F;core ([#501](https://github.com/reearth/reearth/pull/501)) [`8a681d`](https://github.com/reearth/reearth/commit/8a681d)
- Compute features concurrently on reearth&#x2F;core ([#499](https://github.com/reearth/reearth/pull/499)) [`6448ba`](https://github.com/reearth/reearth/commit/6448ba)
- Improve skipping computing process for 3dtiles on reearth&#x2F;core ([#491](https://github.com/reearth/reearth/pull/491)) [`253b58`](https://github.com/reearth/reearth/commit/253b58)
- Improve 3dtiles features calculation on reearth&#x2F;core ([#489](https://github.com/reearth/reearth/pull/489)) [`1204a6`](https://github.com/reearth/reearth/commit/1204a6)
- Improve expression cache strategy on reearth&#x2F;core ([#488](https://github.com/reearth/reearth/pull/488)) [`324e28`](https://github.com/reearth/reearth/commit/324e28)
- Cache AST for evaluator on reearth&#x2F;core ([#473](https://github.com/reearth/reearth/pull/473)) [`da6bb3`](https://github.com/reearth/reearth/commit/da6bb3)
- Improve blink when feature is updated on reearth&#x2F;core ([#429](https://github.com/reearth/reearth/pull/429)) [`c10a67`](https://github.com/reearth/reearth/commit/c10a67)

#### ✨ Refactor

- Replace &quot;team&quot; word  related to Team state with &quot;workspace&quot; ([#607](https://github.com/reearth/reearth/pull/607)) [`fb254b`](https://github.com/reearth/reearth/commit/fb254b)

#### Miscellaneous Tasks

- Remove redundant workflows [`f2685d`](https://github.com/reearth/reearth/commit/f2685d)
- Remove inter-dependency of web and server workflows [`999e73`](https://github.com/reearth/reearth/commit/999e73)
- Change codeowner [`bffa05`](https://github.com/reearth/reearth/commit/bffa05)
- Upgrade eslint [`d55795`](https://github.com/reearth/reearth/commit/d55795)
- Update dependency cesium to v1.104.0 ([#594](https://github.com/reearth/reearth/pull/594)) [`c57839`](https://github.com/reearth/reearth/commit/c57839)
- Fix storybook is not working ([#536](https://github.com/reearth/reearth/pull/536) [`d52124`](https://github.com/reearth/reearth/commit/d52124)
- Rename asset dir to avoid conflicts with backend API endpoints [`1d9455`](https://github.com/reearth/reearth/commit/1d9455)
- Add offline_access auth scope to support refresh tokens with built-in auth server ([#425](https://github.com/reearth/reearth/pull/425) [`2a2af1`](https://github.com/reearth/reearth/commit/2a2af1)
- Upgrade dependencies ([#391](https://github.com/reearth/reearth/pull/391)) [`7280af`](https://github.com/reearth/reearth/commit/7280af)

#### 

- Add updated SECURITY.md [`bc31ce`](https://github.com/reearth/reearth/commit/bc31ce)
- Update to 0.15.0 for the release [`d9b693`](https://github.com/reearth/reearth/commit/d9b693)
- Mono repo — moving reearth-web [`3b1d8d`](https://github.com/reearth/reearth/commit/3b1d8d)
- Chore(server(: upgrade golangci-lint to v1.51 [`9c8714`](https://github.com/reearth/reearth/commit/9c8714)

## 0.14.1 - 2022-12-21

## 0.14.0 - 2022-12-20

## 0.13.0 - 2022-12-06

### Web

### 🔧 Bug Fixes

- Plugin API viewport.query destroys published pages ([#385](https://github.com/reearth/reearth-web/pull/385)) [`de8c01`](https://github.com/reearth/reearth-web/commit/de8c01)

### Web

### 🚀 Features

- Port the Box layer into the reearth&#x2F;core ([#377](https://github.com/reearth/reearth-web/pull/377)) [`f235f1`](https://github.com/reearth/reearth-web/commit/f235f1)
- Support multi feature for geojson ([#380](https://github.com/reearth/reearth-web/pull/380)) [`d1ee59`](https://github.com/reearth/reearth-web/commit/d1ee59)
- Extend plugin API supports get query ([#374](https://github.com/reearth/reearth-web/pull/374)) [`ccae02`](https://github.com/reearth/reearth-web/commit/ccae02)
- Extend plugin API supports communication ([#364](https://github.com/reearth/reearth-web/pull/364)) [`61e67e`](https://github.com/reearth/reearth-web/commit/61e67e)
- Support 2d mode in navigator ([#360](https://github.com/reearth/reearth-web/pull/360)) [`595dd5`](https://github.com/reearth/reearth-web/commit/595dd5)
- Main logic of the new layer system ([#370](https://github.com/reearth/reearth-web/pull/370)) [`0dd63e`](https://github.com/reearth/reearth-web/commit/0dd63e)

### 🔧 Bug Fixes

- Cache logic for feature ([#379](https://github.com/reearth/reearth-web/pull/379)) [`67bc52`](https://github.com/reearth/reearth-web/commit/67bc52)
- Rendering bug in new layer system ([#375](https://github.com/reearth/reearth-web/pull/375)) [`172988`](https://github.com/reearth/reearth-web/commit/172988)

### ⚡️ Performance

- Improve unnecessary loading the globe image ([#378](https://github.com/reearth/reearth-web/pull/378)) [`4abbba`](https://github.com/reearth/reearth-web/commit/4abbba)

### ✨ Refactor

- Move feature context ([#381](https://github.com/reearth/reearth-web/pull/381)) [`7e18b8`](https://github.com/reearth/reearth-web/commit/7e18b8)

### Miscellaneous Tasks

- Upgrade dependencies ([#373](https://github.com/reearth/reearth-web/pull/373)) [`1fcc75`](https://github.com/reearth/reearth-web/commit/1fcc75)

### Web

### 🚀 Features

- Support Cesium Ion terrain ([#331](https://github.com/reearth/reearth-web/pull/331)) [`e0b99a`](https://github.com/reearth/reearth-web/commit/e0b99a)
- Set Cesium Ion default access token via config ([#365](https://github.com/reearth/reearth-web/pull/365)) [`a257b1`](https://github.com/reearth/reearth-web/commit/a257b1)
- Display policy name on workspace title ([#362](https://github.com/reearth/reearth-web/pull/362)) [`c1c632`](https://github.com/reearth/reearth-web/commit/c1c632)
- Editable box ([#357](https://github.com/reearth/reearth-web/pull/357)) [`92a159`](https://github.com/reearth/reearth-web/commit/92a159)
- Extend plugin API supports move widget ([#346](https://github.com/reearth/reearth-web/pull/346)) [`c82825`](https://github.com/reearth/reearth-web/commit/c82825)
- Extend plugin API supports close widget ([#355](https://github.com/reearth/reearth-web/pull/355)) [`d02578`](https://github.com/reearth/reearth-web/commit/d02578)
- Extend plugin API supports get scene inEditor ([#351](https://github.com/reearth/reearth-web/pull/351)) [`ec0b81`](https://github.com/reearth/reearth-web/commit/ec0b81)
- Add clipping box ([#338](https://github.com/reearth/reearth-web/pull/338)) [`af55f1`](https://github.com/reearth/reearth-web/commit/af55f1)
- Extend plugin API event with modalclose popupclose ([#354](https://github.com/reearth/reearth-web/pull/354)) [`9be75a`](https://github.com/reearth/reearth-web/commit/9be75a)
- Extend plugin API supports get location from screen position ([#350](https://github.com/reearth/reearth-web/pull/350)) [`9a826f`](https://github.com/reearth/reearth-web/commit/9a826f)
- Extend plugin API supports get viewport&[#39](https://github.com/reearth/reearth-web/pull/39);s size ([#342](https://github.com/reearth/reearth-web/pull/342)) [`7b268b`](https://github.com/reearth/reearth-web/commit/7b268b)
- Add Re:Earth favicon ([#349](https://github.com/reearth/reearth-web/pull/349)) [`0395d2`](https://github.com/reearth/reearth-web/commit/0395d2)
- Support acquiring locations with terrain ([#343](https://github.com/reearth/reearth-web/pull/343)) [`596543`](https://github.com/reearth/reearth-web/commit/596543)

### 🔧 Bug Fixes

- Policy messages not aligning with policy ([#368](https://github.com/reearth/reearth-web/pull/368)) [`2871ed`](https://github.com/reearth/reearth-web/commit/2871ed)
- Widget align system alignment issue with installed widgets [`e302cc`](https://github.com/reearth/reearth-web/commit/e302cc)
- Rename animation option to withoutAnimation in lookAt ([#361](https://github.com/reearth/reearth-web/pull/361)) [`846a6e`](https://github.com/reearth/reearth-web/commit/846a6e)
- Post message queue doesn&[#39](https://github.com/reearth/reearth-web/pull/39);t work for modal&#x2F;popup ([#359](https://github.com/reearth/reearth-web/pull/359)) [`abb4ed`](https://github.com/reearth/reearth-web/commit/abb4ed)
- Correct flyToGround destination camera ([#356](https://github.com/reearth/reearth-web/pull/356)) [`225758`](https://github.com/reearth/reearth-web/commit/225758)
- Navigator styles ([#353](https://github.com/reearth/reearth-web/pull/353)) [`971323`](https://github.com/reearth/reearth-web/commit/971323)
- Timeline styles ([#352](https://github.com/reearth/reearth-web/pull/352)) [`c76d36`](https://github.com/reearth/reearth-web/commit/c76d36)
- Alignment issues in Widget Align System ([#344](https://github.com/reearth/reearth-web/pull/344)) [`0e12ea`](https://github.com/reearth/reearth-web/commit/0e12ea)
- Cannot input camera altitude less than 500 with camera pane ([#339](https://github.com/reearth/reearth-web/pull/339)) [`76a169`](https://github.com/reearth/reearth-web/commit/76a169)
- Type error from apollo-client [`7dd2b3`](https://github.com/reearth/reearth-web/commit/7dd2b3)
- Fix the camera offset when keep press on zoom to layer ([#335](https://github.com/reearth/reearth-web/pull/335)) [`ccec33`](https://github.com/reearth/reearth-web/commit/ccec33)

### ⚡️ Performance

- Use lodash-es rather than lodash [`731e54`](https://github.com/reearth/reearth-web/commit/731e54)

### 🧪 Testing

- Gql func to e2e reearth page [`593780`](https://github.com/reearth/reearth-web/commit/593780)
- Introduce playwright to run e2e tests ([#336](https://github.com/reearth/reearth-web/pull/336)) [`3af520`](https://github.com/reearth/reearth-web/commit/3af520)

### Miscellaneous Tasks

- Suppress errors output to the console by Icon [`bd9dc5`](https://github.com/reearth/reearth-web/commit/bd9dc5)
- Upgrade cesium to v1.100.0 [`ec05a0`](https://github.com/reearth/reearth-web/commit/ec05a0)
- Upgrade cesium to v1.99 [`be5b22`](https://github.com/reearth/reearth-web/commit/be5b22)
- Upgrade dependencies ([#345](https://github.com/reearth/reearth-web/pull/345)) [`1efe8c`](https://github.com/reearth/reearth-web/commit/1efe8c)
- Remove eslint-plugin-graphql, add eslint-plugin-playwright, refresh yarn.lock [`28c846`](https://github.com/reearth/reearth-web/commit/28c846)
- Use node 16 to avoid storybook build error with node 18 [`64819e`](https://github.com/reearth/reearth-web/commit/64819e)

### Server

#### 🚀 Features

- Remove gsi terrain [`232b78`](https://github.com/reearth/reearth/commit/232b78)
- Suppprt h2c [`ca0626`](https://github.com/reearth/reearth/commit/ca0626)
- Add Cesium Ion and GSI terrain, transfer terrain properties ([#369](https://github.com/reearth/reearth/pull/369)) [`104e59`](https://github.com/reearth/reearth/commit/104e59)

#### 🔧 Bug Fixes

- Enforce policy of layer count on layer group creation correctly [`193c37`](https://github.com/reearth/reearth/commit/193c37)
- Dataset limitation by policies, asset size calc [`e07b78`](https://github.com/reearth/reearth/commit/e07b78)
- Apply default policy to workspaces [`2f7fb9`](https://github.com/reearth/reearth/commit/2f7fb9)

### Misc

#### 🚀 Features

- Policy name, dataset limitation by policies [`b72132`](https://github.com/reearth/reearth/commit/b72132)

#### 🔧 Bug Fixes

- Bugs with incorrectly applied policies ([#377](https://github.com/reearth/reearth/pull/377)) [`67e79f`](https://github.com/reearth/reearth/commit/67e79f)

### 

#### 🚀 Features

- Policy name, dataset limitation by policies [`b72132`](https://github.com/reearth/reearth/commit/b72132)

#### 🔧 Bug Fixes

- Bugs with incorrectly applied policies ([#377](https://github.com/reearth/reearth/pull/377)) [`67e79f`](https://github.com/reearth/reearth/commit/67e79f)

## 0.12.0 - 2022-10-28

### Web

### Miscellaneous Tasks

- Update eslint-config-reearth to 0.2.1 ([#326](https://github.com/reearth/reearth-web/pull/326)) [`25acdd`](https://github.com/reearth/reearth-web/commit/25acdd)

### 🚀 Features

- Add plugin api modal &amp; popup ([#328](https://github.com/reearth/reearth-web/pull/328)) [`27cd7a`](https://github.com/reearth/reearth-web/commit/27cd7a)
- Zoom to layer ([#301](https://github.com/reearth/reearth-web/pull/301)) [`1f5296`](https://github.com/reearth/reearth-web/commit/1f5296)
- Fix layer(marker) extrdue line disapear  ([#330](https://github.com/reearth/reearth-web/pull/330)) [`7c304f`](https://github.com/reearth/reearth-web/commit/7c304f)
- Add option to allow camera to enter the earth&[#39](https://github.com/reearth/reearth-web/pull/39);s surface ([#329](https://github.com/reearth/reearth-web/pull/329)) [`6255ad`](https://github.com/reearth/reearth-web/commit/6255ad)
- Add navigator as a built-in widget ([#323](https://github.com/reearth/reearth-web/pull/323)) [`3befd4`](https://github.com/reearth/reearth-web/commit/3befd4)

### 🔧 Bug Fixes

- Fix the story telling icon size when the layer name is too long  ([#334](https://github.com/reearth/reearth-web/pull/334)) [`c70121`](https://github.com/reearth/reearth-web/commit/c70121)
- Polyfill event target for old Safari [`227d64`](https://github.com/reearth/reearth-web/commit/227d64)
- Zoom to layer functionality ([#332](https://github.com/reearth/reearth-web/pull/332)) [`79b0dd`](https://github.com/reearth/reearth-web/commit/79b0dd)
- Camera popup icon is too small and the icon on storytelling ([#324](https://github.com/reearth/reearth-web/pull/324)) [`e676c3`](https://github.com/reearth/reearth-web/commit/e676c3)

### Miscellaneous Tasks

- Update yarn.lock [`ae4660`](https://github.com/reearth/reearth-web/commit/ae4660)
- Fix wdyr [`d4075a`](https://github.com/reearth/reearth-web/commit/d4075a)
- Update dependency cesium to v1.98.1 ([#325](https://github.com/reearth/reearth-web/pull/325)) [`16e38c`](https://github.com/reearth/reearth-web/commit/16e38c)

### Server

#### 🚀 Features

- Add field to manifest to allow for entering the ground ([#353](https://github.com/reearth/reearth/pull/353)) [`6a56ce`](https://github.com/reearth/reearth/commit/6a56ce)
- Add builtint navigator widget manifest ([#342](https://github.com/reearth/reearth/pull/342)) [`f6834f`](https://github.com/reearth/reearth/commit/f6834f)

#### 🔧 Bug Fixes

- Tiles typo error ([#360](https://github.com/reearth/reearth/pull/360)) [`ddf7d5`](https://github.com/reearth/reearth/commit/ddf7d5)
- Japanese typo error ([#351](https://github.com/reearth/reearth/pull/351)) [`d0292c`](https://github.com/reearth/reearth/commit/d0292c)
- Mongo indexes deleted unexpectedly [`4a323b`](https://github.com/reearth/reearth/commit/4a323b)

#### ⚡️ Performance

- Add mongo indexes [`cee2b2`](https://github.com/reearth/reearth/commit/cee2b2)

#### Miscellaneous Tasks

- Fix make run-db [`944c0d`](https://github.com/reearth/reearth/commit/944c0d)

## 0.11.0 - 2022-10-04

### Web

### 🔧 Bug Fixes

- Installed plugins are not displayed expectedly [`12d546`](https://github.com/reearth/reearth-web/commit/12d546)

### 🚀 Features

- Add 3D OSM building tiles ([#315](https://github.com/reearth/reearth-web/pull/315)) [`2eb89a`](https://github.com/reearth/reearth-web/commit/2eb89a)
- Extend plugin API with camera control ([#311](https://github.com/reearth/reearth-web/pull/311)) [`c1190b`](https://github.com/reearth/reearth-web/commit/c1190b)
- Extend plugin api with captureScreen ([#310](https://github.com/reearth/reearth-web/pull/310)) [`f03022`](https://github.com/reearth/reearth-web/commit/f03022)
- Small update to dashboard UI&#x2F;UX and add marketplace button [`f580e6`](https://github.com/reearth/reearth-web/commit/f580e6)
- Add global modal component and extension location ([#302](https://github.com/reearth/reearth-web/pull/302)) [`7362aa`](https://github.com/reearth/reearth-web/commit/7362aa)
- Add overlay warning to earth editor if browser width is too narrow ([#304](https://github.com/reearth/reearth-web/pull/304)) [`7f5f91`](https://github.com/reearth/reearth-web/commit/7f5f91)
- Add builtin timeline widget ([#285](https://github.com/reearth/reearth-web/pull/285)) [`f774ee`](https://github.com/reearth/reearth-web/commit/f774ee)
- Add plugin settings extension support ([#293](https://github.com/reearth/reearth-web/pull/293)) [`abd1c2`](https://github.com/reearth/reearth-web/commit/abd1c2)

### 🔧 Bug Fixes

- Remove &#x60;show&#x60; props from modal extension ([#321](https://github.com/reearth/reearth-web/pull/321)) [`62eb73`](https://github.com/reearth/reearth-web/commit/62eb73)
- Styles not working as expected in plugin ([#322](https://github.com/reearth/reearth-web/pull/322)) [`21329d`](https://github.com/reearth/reearth-web/commit/21329d)
- Small type error [`21c8bc`](https://github.com/reearth/reearth-web/commit/21c8bc)
- Stop redirect when an error occurs on top page [`236354`](https://github.com/reearth/reearth-web/commit/236354)
- Installed plugins are not correctly listed [`ec305d`](https://github.com/reearth/reearth-web/commit/ec305d)
- Exposed plugin api add layer does not return layer id ([#320](https://github.com/reearth/reearth-web/pull/320)) [`f504d2`](https://github.com/reearth/reearth-web/commit/f504d2)
- Cannot to upgrade marketplace plugins ([#319](https://github.com/reearth/reearth-web/pull/319)) [`444fce`](https://github.com/reearth/reearth-web/commit/444fce)
- Timeline width exceeding browser width ([#316](https://github.com/reearth/reearth-web/pull/316)) [`c6dbb3`](https://github.com/reearth/reearth-web/commit/c6dbb3)
- 3d tile not updating until source type is selected ([#318](https://github.com/reearth/reearth-web/pull/318)) [`49a07b`](https://github.com/reearth/reearth-web/commit/49a07b)
- Wrapper styles in plugin section for plugin extension ([#317](https://github.com/reearth/reearth-web/pull/317)) [`235440`](https://github.com/reearth/reearth-web/commit/235440)
- Timeline speed bug ([#314](https://github.com/reearth/reearth-web/pull/314)) [`984be8`](https://github.com/reearth/reearth-web/commit/984be8)
- Pass extensions to a plugin library extension [`744154`](https://github.com/reearth/reearth-web/commit/744154)
- Print errors when failed to load extensions [`35a63f`](https://github.com/reearth/reearth-web/commit/35a63f)
- Icon button transition happening when undesired ([#313](https://github.com/reearth/reearth-web/pull/313)) [`6a6d98`](https://github.com/reearth/reearth-web/commit/6a6d98)
- Icon size being too small ([#309](https://github.com/reearth/reearth-web/pull/309)) [`3574f0`](https://github.com/reearth/reearth-web/commit/3574f0)
- Change query param used by marketplace to be more concise [`9ba28c`](https://github.com/reearth/reearth-web/commit/9ba28c)
- Dashboard project name not showing ([#307](https://github.com/reearth/reearth-web/pull/307)) [`6a30f5`](https://github.com/reearth/reearth-web/commit/6a30f5)
- Plugin API update event gets called repeatedly, missing hook deps ([#306](https://github.com/reearth/reearth-web/pull/306)) [`47ec24`](https://github.com/reearth/reearth-web/commit/47ec24)
- Update cesium to 1.97.0 [`4d993a`](https://github.com/reearth/reearth-web/commit/4d993a)
- Play button does not work and some unexpected style on timeline widget ([#305](https://github.com/reearth/reearth-web/pull/305)) [`ce29a5`](https://github.com/reearth/reearth-web/commit/ce29a5)
- Development error occurs when updating or DnD layers ([#299](https://github.com/reearth/reearth-web/pull/299)) [`efd079`](https://github.com/reearth/reearth-web/commit/efd079)
- Project creation not creating a scene before earth editor becoming accessible [`a0a03e`](https://github.com/reearth/reearth-web/commit/a0a03e)
- Device settings menu icon getting squished ([#298](https://github.com/reearth/reearth-web/pull/298)) [`91c312`](https://github.com/reearth/reearth-web/commit/91c312)
- Remove unused GraphQL calls [`59f402`](https://github.com/reearth/reearth-web/commit/59f402)

### 🎨 Styling

- Update button and icon UX [`75e6f4`](https://github.com/reearth/reearth-web/commit/75e6f4)

### Miscellaneous Tasks

- Update dependency cesium to v1.97.0 ([#291](https://github.com/reearth/reearth-web/pull/291)) [`dcdf93`](https://github.com/reearth/reearth-web/commit/dcdf93)
- Add remaining props to plugin settings extensions ([#312](https://github.com/reearth/reearth-web/pull/312)) [`893d32`](https://github.com/reearth/reearth-web/commit/893d32)
- Plugin component supports modalContainer and popupContainer props ([#300](https://github.com/reearth/reearth-web/pull/300)) [`fc5f58`](https://github.com/reearth/reearth-web/commit/fc5f58)
- Update cesium to v1.96 ([#303](https://github.com/reearth/reearth-web/pull/303)) [`4fc241`](https://github.com/reearth/reearth-web/commit/4fc241)

### Server

#### 🚀 Features

- Notify plugin donwload to marketplace ([#341](https://github.com/reearth/reearth/pull/341)) [`59f8a8`](https://github.com/reearth/reearth/commit/59f8a8)
- Add 3D OSM building tiles ([#340](https://github.com/reearth/reearth/pull/340)) [`b52132`](https://github.com/reearth/reearth/commit/b52132)
- Allow defining policies to limit functionality on workspaces ([#325](https://github.com/reearth/reearth/pull/325)) [`91ace0`](https://github.com/reearth/reearth/commit/91ace0)

#### 🔧 Bug Fixes

- Prevent API caching [`76405b`](https://github.com/reearth/reearth/commit/76405b)
- Auth server request indexes [`430da0`](https://github.com/reearth/reearth/commit/430da0)
- Marketplace http client bug [`b652c8`](https://github.com/reearth/reearth/commit/b652c8)
- Marketplace http client [`ce982d`](https://github.com/reearth/reearth/commit/ce982d)
- Print marketplace donwload url [`6e1d50`](https://github.com/reearth/reearth/commit/6e1d50)
- Marketplace client init [`752872`](https://github.com/reearth/reearth/commit/752872)
- Support marketplace without auth [`4afe99`](https://github.com/reearth/reearth/commit/4afe99)

#### ✨ Refactor

- Use mongox, update go to 1.19 ([#334](https://github.com/reearth/reearth/pull/334)) [`cfff17`](https://github.com/reearth/reearth/commit/cfff17)
- Separate Team from User, rename Team to Workspace ([#324](https://github.com/reearth/reearth/pull/324)) [`03a94a`](https://github.com/reearth/reearth/commit/03a94a)
- Replace some packages with reearthx ([#322](https://github.com/reearth/reearth/pull/322)) [`3813a9`](https://github.com/reearth/reearth/commit/3813a9)

#### 🧪 Testing

- Add package for e2e tests [`cf7ca5`](https://github.com/reearth/reearth/commit/cf7ca5)

#### Miscellaneous Tasks

- Use reearthx logger [`6b6b21`](https://github.com/reearth/reearth/commit/6b6b21)
- Upgrade mongo to v5 [`9d231d`](https://github.com/reearth/reearth/commit/9d231d)
- Replace package name [`bc1ffe`](https://github.com/reearth/reearth/commit/bc1ffe)

### Misc

#### 🚀 Features

- Installing plugins from marketplace ([#162](https://github.com/reearth/reearth/pull/162)) [`276ef5`](https://github.com/reearth/reearth/commit/276ef5)

#### 🔧 Bug Fixes

- Dataset fails to be loaded [`518f03`](https://github.com/reearth/reearth/commit/518f03)

#### 📖 Documentation

- Update readme [`c8d2ec`](https://github.com/reearth/reearth/commit/c8d2ec)

#### ✨ Refactor

- Use reearthx.authserver ([#335](https://github.com/reearth/reearth/pull/335)) [`83dea5`](https://github.com/reearth/reearth/commit/83dea5)

#### Miscellaneous Tasks

- Merge &#x60;reearth&#x2F;reearth-backend&#x60; ([#318](https://github.com/reearth/reearth/pull/318)) [`98514b`](https://github.com/reearth/reearth/commit/98514b)

### 

#### 🚀 Features

- Installing plugins from marketplace ([#162](https://github.com/reearth/reearth/pull/162)) [`276ef5`](https://github.com/reearth/reearth/commit/276ef5)

#### 🔧 Bug Fixes

- Dataset fails to be loaded [`518f03`](https://github.com/reearth/reearth/commit/518f03)

#### 📖 Documentation

- Update readme [`c8d2ec`](https://github.com/reearth/reearth/commit/c8d2ec)

#### ✨ Refactor

- Use reearthx.authserver ([#335](https://github.com/reearth/reearth/pull/335)) [`83dea5`](https://github.com/reearth/reearth/commit/83dea5)

#### Miscellaneous Tasks

- Merge &#x60;reearth&#x2F;reearth-backend&#x60; ([#318](https://github.com/reearth/reearth/pull/318)) [`98514b`](https://github.com/reearth/reearth/commit/98514b)

## 0.10.0 - 2022-08-10

### reearth-web

#### 🚀 Features

- Add mouse events to plugin API ([#280](https://github.com/reearth/reearth-web/pull/280)) [`9445f0`](https://github.com/reearth/reearth-web/commit/9445f0)

#### 🔧 Bug Fixes

- Select not working after pinch event on ipad ([#290](https://github.com/reearth/reearth-web/pull/290)) [`821504`](https://github.com/reearth/reearth-web/commit/821504)
- Translation for modal buttons [`7eead9`](https://github.com/reearth/reearth-web/commit/7eead9)
- Plugin widget&[#39](https://github.com/reearth/reearth-web/pull/39);s width using iframe&[#39](https://github.com/reearth/reearth-web/pull/39);s default ([#283](https://github.com/reearth/reearth-web/pull/283)) [`572da0`](https://github.com/reearth/reearth-web/commit/572da0)
- Pointer events issues around widgets ([#279](https://github.com/reearth/reearth-web/pull/279)) [`219ea4`](https://github.com/reearth/reearth-web/commit/219ea4)

#### 🎨 Styling

- Fix icons of plugin install buttons ([#289](https://github.com/reearth/reearth-web/pull/289)) [`af7a1b`](https://github.com/reearth/reearth-web/commit/af7a1b)

#### 🧪 Testing

- Introduce vitest ([#284](https://github.com/reearth/reearth-web/pull/284)) [`2152e0`](https://github.com/reearth/reearth-web/commit/2152e0)

#### Miscellaneous Tasks

- Migrate to Vite, upgrade Cypress to v10 ([#287](https://github.com/reearth/reearth-web/pull/287)) [`50f2b6`](https://github.com/reearth/reearth-web/commit/50f2b6)
- Simplify ESLint config ([#282](https://github.com/reearth/reearth-web/pull/282)) [`b3570a`](https://github.com/reearth/reearth-web/commit/b3570a)
- Upgrade resium to v1.15.0 ([#281](https://github.com/reearth/reearth-web/pull/281)) [`bd3968`](https://github.com/reearth/reearth-web/commit/bd3968)
- Cosme changelog [`05084e`](https://github.com/reearth/reearth-web/commit/05084e)
- Fix changelog [`48de86`](https://github.com/reearth/reearth-web/commit/48de86)

### reearth-backend

#### 🚀 Features

- Configurable server host [`61b03a`](https://github.com/reearth/reearth-backend/commit/61b03a)

#### Miscellaneous Tasks

- Add new frontend endpoint (for Vite@3) [`70fed0`](https://github.com/reearth/reearth-backend/commit/70fed0)
- Fix changelog [skip ci] [`895a64`](https://github.com/reearth/reearth-backend/commit/895a64)

## 0.9.0 - 2022-07-20

### reearth-web

#### 🚀 Features

- Plugin API to add layers ([#258](https://github.com/reearth/reearth-web/pull/258)) [`6468e2`](https://github.com/reearth/reearth-web/commit/6468e2)
- Change layer indicators from preset list ([#245](https://github.com/reearth/reearth-web/pull/245)) [`db185e`](https://github.com/reearth/reearth-web/commit/db185e)

#### 🔧 Bug Fixes

- Some menu not displayed at sidebar in proejct setting page [`7c0705`](https://github.com/reearth/reearth-web/commit/7c0705)
- Nothing displayed at project setting page when there are many projects [`0a6744`](https://github.com/reearth/reearth-web/commit/0a6744)
- Plugins do not work as expected, update quickjs-emscripten ([#276](https://github.com/reearth/reearth-web/pull/276)) [`9336e6`](https://github.com/reearth/reearth-web/commit/9336e6)
- Plugin editor changes do not take effect until run button is clicked ([#274](https://github.com/reearth/reearth-web/pull/274)) [`39fdb2`](https://github.com/reearth/reearth-web/commit/39fdb2)
- Storytelling widget does not get layers&[#39](https://github.com/reearth/reearth-web/pull/39); title ([#273](https://github.com/reearth/reearth-web/pull/273)) [`5ff72b`](https://github.com/reearth/reearth-web/commit/5ff72b)
- Dataset icon not showing in layer list ([#275](https://github.com/reearth/reearth-web/pull/275)) [`8dbc88`](https://github.com/reearth/reearth-web/commit/8dbc88)
- Show full camera values in camera property field popup ([#270](https://github.com/reearth/reearth-web/pull/270)) [`7d3eac`](https://github.com/reearth/reearth-web/commit/7d3eac)
- Plugin dimensions and iframe issues ([#271](https://github.com/reearth/reearth-web/pull/271)) [`f3a52a`](https://github.com/reearth/reearth-web/commit/f3a52a)
- Camera jump not working ([#269](https://github.com/reearth/reearth-web/pull/269)) [`48bbfe`](https://github.com/reearth/reearth-web/commit/48bbfe)
- Layer select state not update properly ([#268](https://github.com/reearth/reearth-web/pull/268)) [`5f7c69`](https://github.com/reearth/reearth-web/commit/5f7c69)
- Unselect layer not work properly ([#266](https://github.com/reearth/reearth-web/pull/266)) [`eb41da`](https://github.com/reearth/reearth-web/commit/eb41da)
- Layer drag and drop does not work with indicators ([#265](https://github.com/reearth/reearth-web/pull/265)) [`12ae04`](https://github.com/reearth/reearth-web/commit/12ae04)
- Testing-library react 18 warnings ([#263](https://github.com/reearth/reearth-web/pull/263)) [`4c9076`](https://github.com/reearth/reearth-web/commit/4c9076)
- Auto fetch more items in dashboard page , project list , dataset page for big screens ([#255](https://github.com/reearth/reearth-web/pull/255)) [`fb8bf9`](https://github.com/reearth/reearth-web/commit/fb8bf9)
- Asset modal flushes when camera limiter is enabled ([#261](https://github.com/reearth/reearth-web/pull/261)) [`204629`](https://github.com/reearth/reearth-web/commit/204629)
- Not being able to override an image from the asset modal ([#260](https://github.com/reearth/reearth-web/pull/260)) [`1d3c3f`](https://github.com/reearth/reearth-web/commit/1d3c3f)
- Layers pane does not update after move layer or create folder  ([#259](https://github.com/reearth/reearth-web/pull/259)) [`336d98`](https://github.com/reearth/reearth-web/commit/336d98)
- Cesium flashes on camera change ([#257](https://github.com/reearth/reearth-web/pull/257)) [`ad2c0e`](https://github.com/reearth/reearth-web/commit/ad2c0e)
- Router typos ([#252](https://github.com/reearth/reearth-web/pull/252)) [`19fcb6`](https://github.com/reearth/reearth-web/commit/19fcb6)
- Dataset page showing errors on page refreshing  ([#253](https://github.com/reearth/reearth-web/pull/253)) [`3f48e9`](https://github.com/reearth/reearth-web/commit/3f48e9)

#### 🧪 Testing

- Fix test coverage target ([#272](https://github.com/reearth/reearth-web/pull/272)) [`b9db10`](https://github.com/reearth/reearth-web/commit/b9db10)

#### Miscellaneous Tasks

- Update dependency cesium to ^1.95.0 ([#262](https://github.com/reearth/reearth-web/pull/262)) [`845e2a`](https://github.com/reearth/reearth-web/commit/845e2a)
- Upgrade cesium [`363071`](https://github.com/reearth/reearth-web/commit/363071)
- Upgrade to React 18 and switch to React Router ([#234](https://github.com/reearth/reearth-web/pull/234)) [`b0e8e6`](https://github.com/reearth/reearth-web/commit/b0e8e6)

### reearth-backend

#### 🚀 Features

- Change layer indicators from preset list from backend side ([#158](https://github.com/reearth/reearth-backend/pull/158)) [`0267f1`](https://github.com/reearth/reearth-backend/commit/0267f1)

#### 🔧 Bug Fixes

- Property fields in a property list cannot be removed ([#160](https://github.com/reearth/reearth-backend/pull/160)) [`358237`](https://github.com/reearth/reearth-backend/commit/358237)

#### 🧪 Testing

- Unit test for mongo auth request repo ([#159](https://github.com/reearth/reearth-backend/pull/159)) [`5afc81`](https://github.com/reearth/reearth-backend/commit/5afc81)

#### Miscellaneous Tasks

- Update Makefile to remove unused targets [`67780b`](https://github.com/reearth/reearth-backend/commit/67780b)


## 0.8.0 - 2022-06-17

### reearth-web

#### 🚀 Features

- Add a basic timeline UI ([#232](https://github.com/reearth/reearth-web/pull/232)) [`fc9732`](https://github.com/reearth/reearth-web/commit/fc9732)
- Add infinite scroll for project lists and datasets in dashboard and setting pages ([#225](https://github.com/reearth/reearth-web/pull/225)) [`28d377`](https://github.com/reearth/reearth-web/commit/28d377)
- Adapt camera field to support 2d mode ([#233](https://github.com/reearth/reearth-web/pull/233)) [`172de5`](https://github.com/reearth/reearth-web/commit/172de5)
- Add scene property overriding to Re:Earth API ([#224](https://github.com/reearth/reearth-web/pull/224)) [`b07603`](https://github.com/reearth/reearth-web/commit/b07603)

#### 🔧 Bug Fixes

- Some plugin APIs were missing ([#248](https://github.com/reearth/reearth-web/pull/248)) [`c83262`](https://github.com/reearth/reearth-web/commit/c83262)
- Slight shift when capture a new position ([#246](https://github.com/reearth/reearth-web/pull/246)) [`182406`](https://github.com/reearth/reearth-web/commit/182406)
- Dataset counts are displayed incorrectly in dataset pane ([#235](https://github.com/reearth/reearth-web/pull/235)) [`45a0c8`](https://github.com/reearth/reearth-web/commit/45a0c8)
- Labeling hidden by marker symbol ([#238](https://github.com/reearth/reearth-web/pull/238)) [`99c378`](https://github.com/reearth/reearth-web/commit/99c378)
- Vertical position style in infobox image block ([#236](https://github.com/reearth/reearth-web/pull/236)) [`647cf8`](https://github.com/reearth/reearth-web/commit/647cf8)
- Unexpected values for theme and lang props of extension components [`723486`](https://github.com/reearth/reearth-web/commit/723486)
- Wait until all extensions are loaded [`dfe2aa`](https://github.com/reearth/reearth-web/commit/dfe2aa)
- Iframe not correctly sizing to plugin ([#230](https://github.com/reearth/reearth-web/pull/230)) [`500ce8`](https://github.com/reearth/reearth-web/commit/500ce8)
- Plugin API cameramove event is not emitted in published pages ([#227](https://github.com/reearth/reearth-web/pull/227)) [`7a11b3`](https://github.com/reearth/reearth-web/commit/7a11b3)

#### ✨ Refactor

- Migrate to react-intl from react-i18next ([#240](https://github.com/reearth/reearth-web/pull/240)) [`404743`](https://github.com/reearth/reearth-web/commit/404743)

#### 🧪 Testing

- Disable util/raf tests that do not always succeed [`45a450`](https://github.com/reearth/reearth-web/commit/45a450)
- Fix unit test for utils/raf [`a060d9`](https://github.com/reearth/reearth-web/commit/a060d9)
- Fix Cypress login test fails ([#241](https://github.com/reearth/reearth-web/pull/241)) [`a5dbfb`](https://github.com/reearth/reearth-web/commit/a5dbfb)

#### Miscellaneous Tasks

- Upgrade dependency cesium-dnd to 1.1.0 ([#244](https://github.com/reearth/reearth-web/pull/244)) [`ba6b51`](https://github.com/reearth/reearth-web/commit/ba6b51)
- Fix typos [`f98005`](https://github.com/reearth/reearth-web/commit/f98005)
- Update config so extensionUrls can be declared in .env file for local development ([#237](https://github.com/reearth/reearth-web/pull/237)) [`545b9e`](https://github.com/reearth/reearth-web/commit/545b9e)

### reearth-backend

#### 🚀 Features

- Add totalCount field to DatasetSchema type of GraphQL schema ([#154](https://github.com/reearth/reearth-backend/pull/154)) [`ab6334`](https://github.com/reearth/reearth-backend/commit/ab6334)
- Add timeline settings to scene property ([#153](https://github.com/reearth/reearth-backend/pull/153)) [`602ec0`](https://github.com/reearth/reearth-backend/commit/602ec0)

#### 🔧 Bug Fixes

- Assets are not saved when files are uploaded ([#155](https://github.com/reearth/reearth-backend/pull/155)) [`e444e4`](https://github.com/reearth/reearth-backend/commit/e444e4)

#### ✨ Refactor

- Declarative description of use case structure (asset only) ([#151](https://github.com/reearth/reearth-backend/pull/151)) [`c6e98c`](https://github.com/reearth/reearth-backend/commit/c6e98c)

#### Miscellaneous Tasks

- Update go modules ([#150](https://github.com/reearth/reearth-backend/pull/150)) [`6372bc`](https://github.com/reearth/reearth-backend/commit/6372bc)

## 0.7.0 - 2022-05-17

### reearth-web

#### 🚀 Features

- Implementation of the avatar feature in workspaces screens ([#206](https://github.com/reearth/reearth-web/pull/206)) [`42d7aa`](https://github.com/reearth/reearth-web/commit/42d7aa)
- Update placehoder for color field ([#215](https://github.com/reearth/reearth-web/pull/215)) [`c6c6e3`](https://github.com/reearth/reearth-web/commit/c6c6e3)
- Add opacity field to map tiles ([#220](https://github.com/reearth/reearth-web/pull/220)) [`006a8d`](https://github.com/reearth/reearth-web/commit/006a8d)

#### 🔧 Bug Fixes

- Dropdown styles in right panel break when selected item's name is too long [`9a5993`](https://github.com/reearth/reearth-web/commit/9a5993)
- Dashboard not updating on project creation [`4b5478`](https://github.com/reearth/reearth-web/commit/4b5478)
- Query names in refetchQueries not updated ([#222](https://github.com/reearth/reearth-web/pull/222)) [`711712`](https://github.com/reearth/reearth-web/commit/711712)
- Published page uses GraphQL and reports errors [`3e3e45`](https://github.com/reearth/reearth-web/commit/3e3e45)

#### ✨ Refactor

- Queries/mutation code into a single directory ([#208](https://github.com/reearth/reearth-web/pull/208)) [`2afc16`](https://github.com/reearth/reearth-web/commit/2afc16)
- Property, scene, tag, user, widget and workspace gql query files ([#221](https://github.com/reearth/reearth-web/pull/221)) [`3bf421`](https://github.com/reearth/reearth-web/commit/3bf421)

#### Miscellaneous Tasks

- Introduce i18next ([#212](https://github.com/reearth/reearth-web/pull/212)) [`0ac0c2`](https://github.com/reearth/reearth-web/commit/0ac0c2)
- Add reference to style guide in README [`e29024`](https://github.com/reearth/reearth-web/commit/e29024)
- Add useT hook to i18n ([#223](https://github.com/reearth/reearth-web/pull/223)) [`b96177`](https://github.com/reearth/reearth-web/commit/b96177)
- Update dependency cesium to ^1.93.0 ([#216](https://github.com/reearth/reearth-web/pull/216)) [`06b563`](https://github.com/reearth/reearth-web/commit/06b563)
- Update all dependencies ([#226](https://github.com/reearth/reearth-web/pull/226)) [`36fb79`](https://github.com/reearth/reearth-web/commit/36fb79)

#### Refactor

- Clean gql pt1 asset ([#217](https://github.com/reearth/reearth-web/pull/217)) [`b88a8c`](https://github.com/reearth/reearth-web/commit/b88a8c)
- Cluster, dataset, infobox, layer, plugin and project gql query files ([#219](https://github.com/reearth/reearth-web/pull/219)) [`e4dae9`](https://github.com/reearth/reearth-web/commit/e4dae9)

### reearth-backend

#### 🚀 Features

- Add an opacity slider to map tiles ([#138](https://github.com/reearth/reearth-backend/pull/138)) [`4f72b8`](https://github.com/reearth/reearth-backend/commit/4f72b8)

#### 🔧 Bug Fixes

- Signup api requires password field [`a79376`](https://github.com/reearth/reearth-backend/commit/a79376)
- "$in needs an array" error from mongo FindByIDs ([#142](https://github.com/reearth/reearth-backend/pull/142)) [`58e1b0`](https://github.com/reearth/reearth-backend/commit/58e1b0)
- Name field is available again in signup api ([#144](https://github.com/reearth/reearth-backend/pull/144)) [`651852`](https://github.com/reearth/reearth-backend/commit/651852)

#### ✨ Refactor

- Retry mongo lock ([#145](https://github.com/reearth/reearth-backend/pull/145)) [`ddaeaa`](https://github.com/reearth/reearth-backend/commit/ddaeaa)

#### 🧪 Testing

- Add Mongo Asset's [`FindByID`](https://github.com/reearth/reearth-backend/commit/FindByID) unit testing ([#139](https://github.com/reearth/reearth-backend/pull/139)) [`35f9db`](https://github.com/reearth/reearth-backend/commit/35f9db)
- Refactor mongo connect helper function [`751e66`](https://github.com/reearth/reearth-backend/commit/751e66)
- Util.SyncMap.Range test sometimes fails ([#143](https://github.com/reearth/reearth-backend/pull/143)) [`c2b969`](https://github.com/reearth/reearth-backend/commit/c2b969)

#### Miscellaneous Tasks

- Typo [`secrit`](https://github.com/reearth/reearth-backend/commit/secrit) on env example ([#137](https://github.com/reearth/reearth-backend/pull/137)) [`2c0220`](https://github.com/reearth/reearth-backend/commit/2c0220)
- Update the go modules ([#146](https://github.com/reearth/reearth-backend/pull/146)) [`89009b`](https://github.com/reearth/reearth-backend/commit/89009b)

## 0.6.1 - 2022-04-20

### reearth-web

#### 🚀 Features

- Extend project publish settings and dataset import modal functionality through extension API ([#200](https://github.com/reearth/reearth-web/pull/200)) [`96aa56`](https://github.com/reearth/reearth-web/commit/96aa56)

#### 🔧 Bug Fixes

- Redirect after esc button in any setting page ([#193](https://github.com/reearth/reearth-web/pull/193)) [`c8ec35`](https://github.com/reearth/reearth-web/commit/c8ec35)

#### Miscellaneous Tasks

- Follow GraphQL schema updates ([#209](https://github.com/reearth/reearth-web/pull/209)) [`8d9e75`](https://github.com/reearth/reearth-web/commit/8d9e75)
- Update all dependencies ([#210](https://github.com/reearth/reearth-web/pull/210)) [`c22b7a`](https://github.com/reearth/reearth-web/commit/c22b7a)

### reearth-backend

#### 🔧 Bug Fixes

- Renovate bot not running on schedule ([#136](https://github.com/reearth/reearth-backend/pull/136)) [`82843f`](https://github.com/reearth/reearth-backend/commit/82843f)
- Aud was changed and jwt could not be validated correctly [`985100`](https://github.com/reearth/reearth-backend/commit/985100)
- Auth audiences were unintentionally required [`7ec76a`](https://github.com/reearth/reearth-backend/commit/7ec76a)

#### ✨ Refactor

- Introduce generics, reorganize GraphQL schema ([#135](https://github.com/reearth/reearth-backend/pull/135)) [`04a098`](https://github.com/reearth/reearth-backend/commit/04a098)

#### Miscellaneous Tasks

- Update dependencies ([#134](https://github.com/reearth/reearth-backend/pull/134)) [`1b9b6b`](https://github.com/reearth/reearth-backend/commit/1b9b6b)

## 0.6.0 - 2022-04-08

### reearth-web

#### 🚀 Features

- Add a plugin API to resize iframe by plugins ([#181](https://github.com/reearth/reearth-web/pull/181)) [`7c1019`](https://github.com/reearth/reearth-web/commit/7c1019)
- Authentication ([#121](https://github.com/reearth/reearth-web/pull/121)) [`b63018`](https://github.com/reearth/reearth-web/commit/b63018)
- Infinite scroll on assets ([#130](https://github.com/reearth/reearth-web/pull/130)) [`11f2f2`](https://github.com/reearth/reearth-web/commit/11f2f2)
- Basic plugin editor ([#184](https://github.com/reearth/reearth-web/pull/184)) [`1c4e09`](https://github.com/reearth/reearth-web/commit/1c4e09)

#### 🔧 Bug Fixes

- Unable to type RGBA values ([#180](https://github.com/reearth/reearth-web/pull/180)) [`f7345c`](https://github.com/reearth/reearth-web/commit/f7345c)
- Small height of block plugins [`8070a3`](https://github.com/reearth/reearth-web/commit/8070a3)
- Button widget squishing its text & infobox mask click away ([#185](https://github.com/reearth/reearth-web/pull/185)) [`ac7ef0`](https://github.com/reearth/reearth-web/commit/ac7ef0)
- Cannot select layers that activate infobox mask ([#186](https://github.com/reearth/reearth-web/pull/186)) [`d824b6`](https://github.com/reearth/reearth-web/commit/d824b6)
- Display error messages from auth server ([#187](https://github.com/reearth/reearth-web/pull/187)) [`e19fab`](https://github.com/reearth/reearth-web/commit/e19fab)
- Duplicate asset results ([#188](https://github.com/reearth/reearth-web/pull/188)) [`b3eb7f`](https://github.com/reearth/reearth-web/commit/b3eb7f)
- Workspace name cannot be changed, error displayed when deleting assets ([#189](https://github.com/reearth/reearth-web/pull/189)) [`a99cf3`](https://github.com/reearth/reearth-web/commit/a99cf3)
- Multiple assets in infinite scroll and datasets not showing in DatasetPane  ([#192](https://github.com/reearth/reearth-web/pull/192)) [`6f5c93`](https://github.com/reearth/reearth-web/commit/6f5c93)
- Asset modal showing only image-based assets ([#196](https://github.com/reearth/reearth-web/pull/196)) [`83a6bf`](https://github.com/reearth/reearth-web/commit/83a6bf)
- Screen becomes inoperable when errors occur in sign up [`820a04`](https://github.com/reearth/reearth-web/commit/820a04)
- Add missing translations [`a4c237`](https://github.com/reearth/reearth-web/commit/a4c237)

#### Miscellaneous Tasks

- Update dependency cesium to ^1.91.0 ([#182](https://github.com/reearth/reearth-web/pull/182)) [`603a5c`](https://github.com/reearth/reearth-web/commit/603a5c)
- Set default auth config to start app with zero configuration ([#191](https://github.com/reearth/reearth-web/pull/191)) [`d5a2aa`](https://github.com/reearth/reearth-web/commit/d5a2aa)

### reearth-backend

#### 🚀 Features

- Authentication system ([#108](https://github.com/reearth/reearth-backend/pull/108)) [`b89c32`](https://github.com/reearth/reearth-backend/commit/b89c32)
- Default mailer that outputs mails into stdout [`aab26c`](https://github.com/reearth/reearth-backend/commit/aab26c)
- Assets filtering & pagination ([#81](https://github.com/reearth/reearth-backend/pull/81)) [`739943`](https://github.com/reearth/reearth-backend/commit/739943)
- Support sign up with information provided by OIDC providers ([#130](https://github.com/reearth/reearth-backend/pull/130)) [`fef60e`](https://github.com/reearth/reearth-backend/commit/fef60e)

#### 🔧 Bug Fixes

- Load auth client domain from config ([#124](https://github.com/reearth/reearth-backend/pull/124)) [`9bde8a`](https://github.com/reearth/reearth-backend/commit/9bde8a)
- Signup fails when password is not set [`27c2f0`](https://github.com/reearth/reearth-backend/commit/27c2f0)
- Logger panics [`d1e3a8`](https://github.com/reearth/reearth-backend/commit/d1e3a8)
- Set auth server dev mode automatically [`83a66a`](https://github.com/reearth/reearth-backend/commit/83a66a)
- Auth server bugs and auth client bugs ([#125](https://github.com/reearth/reearth-backend/pull/125)) [`ce2309`](https://github.com/reearth/reearth-backend/commit/ce2309)
- Auth0 setting is not used by JWT verification middleware [`232e75`](https://github.com/reearth/reearth-backend/commit/232e75)
- Invalid mongo queries of pagination [`7caf68`](https://github.com/reearth/reearth-backend/commit/7caf68)
- Auth config not loaded expectedly [`570fe7`](https://github.com/reearth/reearth-backend/commit/570fe7)
- Users cannot creates a new team and scene [`5df25f`](https://github.com/reearth/reearth-backend/commit/5df25f)
- Auth server certificate is not saved as pem format [`982a71`](https://github.com/reearth/reearth-backend/commit/982a71)
- Repo filters are not merged expectedly [`f4cc3f`](https://github.com/reearth/reearth-backend/commit/f4cc3f)
- Auth is no longer required for GraphQL endpoint [`58a6d1`](https://github.com/reearth/reearth-backend/commit/58a6d1)
- Rename auth srv default client ID ([#128](https://github.com/reearth/reearth-backend/pull/128)) [`89adc3`](https://github.com/reearth/reearth-backend/commit/89adc3)
- Signup API is disabled when auth server is disabled, users and auth requests in mongo cannot be deleted ([#132](https://github.com/reearth/reearth-backend/pull/132)) [`47be6a`](https://github.com/reearth/reearth-backend/commit/47be6a)
- Auth to work with zero config ([#131](https://github.com/reearth/reearth-backend/pull/131)) [`3cbb45`](https://github.com/reearth/reearth-backend/commit/3cbb45)
- Property.SchemaListMap.List test fails [`3e6dff`](https://github.com/reearth/reearth-backend/commit/3e6dff)
- Errors when auth srv domain is not specified [`10691a`](https://github.com/reearth/reearth-backend/commit/10691a)
- Errors when auth srv domain is not specified [`648073`](https://github.com/reearth/reearth-backend/commit/648073)
- Login redirect does not work [`cb6ca4`](https://github.com/reearth/reearth-backend/commit/cb6ca4)
- Enable auth srv dev mode when no domain is specified [`0c0e28`](https://github.com/reearth/reearth-backend/commit/0c0e28)
- Add a trailing slash to jwt audiences [`e96f78`](https://github.com/reearth/reearth-backend/commit/e96f78)
- Allow separate auth server ui domain [`0ce79f`](https://github.com/reearth/reearth-backend/commit/0ce79f)

#### ⚡️ Performance

- Reduce database queries to obtain scene IDs ([#119](https://github.com/reearth/reearth-backend/pull/119)) [`784332`](https://github.com/reearth/reearth-backend/commit/784332)

#### ✨ Refactor

- Remove filter args from repos to prevent implementation errors in the use case layer ([#122](https://github.com/reearth/reearth-backend/pull/122)) [`82cf28`](https://github.com/reearth/reearth-backend/commit/82cf28)
- Http api to export layers [`3f2582`](https://github.com/reearth/reearth-backend/commit/3f2582)

#### Miscellaneous Tasks

- Update dependencies ([#117](https://github.com/reearth/reearth-backend/pull/117)) [`d1a38e`](https://github.com/reearth/reearth-backend/commit/d1a38e)
- Update docker-compose config [`83f9b1`](https://github.com/reearth/reearth-backend/commit/83f9b1)
- Add log for GraphQL Playground endpoint ([#133](https://github.com/reearth/reearth-backend/pull/133)) [`adeda4`](https://github.com/reearth/reearth-backend/commit/adeda4)

## 0.5.0 - 2022-02-24

### reearth-web

#### 🚀 Features

- Allowing widget and block plugins to resize when they are expandable ([#170](https://github.com/reearth/reearth-web/pull/170)) [`4fdf5f`](https://github.com/reearth/reearth-web/commit/4fdf5f)
- Plugin APIs to get camera viewport and layers in the viewport ([#165](https://github.com/reearth/reearth-web/pull/165)) [`f1f95a`](https://github.com/reearth/reearth-web/commit/f1f95a)
- Improving the Infobox style  ([#176](https://github.com/reearth/reearth-web/pull/176)) [`f1ddda`](https://github.com/reearth/reearth-web/commit/f1ddda)

#### 🔧 Bug Fixes

- Plugin blocks cannot be deleted ([#164](https://github.com/reearth/reearth-web/pull/164)) [`a4f17f`](https://github.com/reearth/reearth-web/commit/a4f17f)
- Support tree-structured layers and tags in published pages ([#168](https://github.com/reearth/reearth-web/pull/168)) [`17d968`](https://github.com/reearth/reearth-web/commit/17d968)
- Workspace settings does not refresh ([#167](https://github.com/reearth/reearth-web/pull/167)) [`0f3654`](https://github.com/reearth/reearth-web/commit/0f3654)
- Plugin layersInViewport API returns errors for layers that have no location fields [`e52b44`](https://github.com/reearth/reearth-web/commit/e52b44)

#### ✨ Refactor

- Format codes [`219ac6`](https://github.com/reearth/reearth-web/commit/219ac6)
- Format codes [`4e5b61`](https://github.com/reearth/reearth-web/commit/4e5b61)

#### Miscellaneous Tasks

- Upgrade dependencies ([#175](https://github.com/reearth/reearth-web/pull/175)) [`dba959`](https://github.com/reearth/reearth-web/commit/dba959)

### reearth-backend

#### 🚀 Features

- Implement property.Diff and plugin/manifest.Diff ([#107](https://github.com/reearth/reearth-backend/pull/107)) [`700269`](https://github.com/reearth/reearth-backend/commit/700269)
- Support 3rd party plugin translation ([#109](https://github.com/reearth/reearth-backend/pull/109)) [`67a618`](https://github.com/reearth/reearth-backend/commit/67a618)
- Improve the Infobox style (manifest) ([#110](https://github.com/reearth/reearth-backend/pull/110)) [`7aebcd`](https://github.com/reearth/reearth-backend/commit/7aebcd)
- Overwrite installation of new plug-ins without removing (automatic property migration) ([#113](https://github.com/reearth/reearth-backend/pull/113)) [`2dc192`](https://github.com/reearth/reearth-backend/commit/2dc192)
- Update infobox style fields ([#115](https://github.com/reearth/reearth-backend/pull/115)) [`608436`](https://github.com/reearth/reearth-backend/commit/608436)

#### 🔧 Bug Fixes

- Scene exporter should export layers and tags while maintaining the tree structure ([#104](https://github.com/reearth/reearth-backend/pull/104)) [`805d78`](https://github.com/reearth/reearth-backend/commit/805d78)
- Property field in groups in list cannot be updated correctly [`5009c5`](https://github.com/reearth/reearth-backend/commit/5009c5)
- Scenes and properties are not updated properly when plugin is updated [`861c4b`](https://github.com/reearth/reearth-backend/commit/861c4b)
- Scene widgets and blocks are not update properly when plugin is updated [`f66f9a`](https://github.com/reearth/reearth-backend/commit/f66f9a)

#### ✨ Refactor

- Graphql resolvers ([#105](https://github.com/reearth/reearth-backend/pull/105)) [`01a4e6`](https://github.com/reearth/reearth-backend/commit/01a4e6)

#### Miscellaneous Tasks

- Update all dependencies ([#111](https://github.com/reearth/reearth-backend/pull/111)) [`173881`](https://github.com/reearth/reearth-backend/commit/173881)
- Increase batch size of db migration [ci skip] [`fbbca4`](https://github.com/reearth/reearth-backend/commit/fbbca4)

## 0.4.0 - 2022-01-27

### reearth-web

#### 🚀 Features

- Add "clamp to filed" option to file primitive ([#155](https://github.com/reearth/reearth-web/pull/155)) [`2e83ba`](https://github.com/reearth/reearth-web/commit/2e83ba)
- Infobox padding ([#158](https://github.com/reearth/reearth-web/pull/158)) [`90084f`](https://github.com/reearth/reearth-web/commit/90084f)
- Support tags in plugin API ([#153](https://github.com/reearth/reearth-web/pull/153)) [`1031c5`](https://github.com/reearth/reearth-web/commit/1031c5)

#### 🔧 Bug Fixes

- Enable to select blocks of plugins ([#162](https://github.com/reearth/reearth-web/pull/162)) [`458402`](https://github.com/reearth/reearth-web/commit/458402)
- Cesium Ion acces token is not set expectedly ([#160](https://github.com/reearth/reearth-web/pull/160)) [`e8e183`](https://github.com/reearth/reearth-web/commit/e8e183)
- Cluster styling issue ([#161](https://github.com/reearth/reearth-web/pull/161)) [`c78872`](https://github.com/reearth/reearth-web/commit/c78872)
- Clusters and layers are not displayed correctly [`4fc124`](https://github.com/reearth/reearth-web/commit/4fc124)
- Type error [`b01bc7`](https://github.com/reearth/reearth-web/commit/b01bc7)
- The style of infobox block dropdown list is broken ([#163](https://github.com/reearth/reearth-web/pull/163)) [`6e02a9`](https://github.com/reearth/reearth-web/commit/6e02a9)
- Plugin blocks protrude from the infobox [`6cf0d3`](https://github.com/reearth/reearth-web/commit/6cf0d3)

#### ✨ Refactor

- Layer clustering feature ([#157](https://github.com/reearth/reearth-web/pull/157)) [`db6e6c`](https://github.com/reearth/reearth-web/commit/db6e6c)
- Camera limiter ([#149](https://github.com/reearth/reearth-web/pull/149)) [`105428`](https://github.com/reearth/reearth-web/commit/105428)
- Layer clustering feature (GraphQL) ([#159](https://github.com/reearth/reearth-web/pull/159)) [`4365b8`](https://github.com/reearth/reearth-web/commit/4365b8)

### reearth-backend

#### 🚀 Features

- Add "clamp to ground" option to file primitive ([#95](https://github.com/reearth/reearth-backend/pull/95)) [`559194`](https://github.com/reearth/reearth-backend/commit/559194)
- Infobox and text block padding ([#100](https://github.com/reearth/reearth-backend/pull/100)) [`ddd0db`](https://github.com/reearth/reearth-backend/commit/ddd0db)

#### ⚡️ Performance

- Add indexes of mongo collections ([#98](https://github.com/reearth/reearth-backend/pull/98)) [`691cb7`](https://github.com/reearth/reearth-backend/commit/691cb7)

#### ✨ Refactor

- Pkg/id, use ID aliases, move JSON schemas ([#97](https://github.com/reearth/reearth-backend/pull/97)) [`1265ac`](https://github.com/reearth/reearth-backend/commit/1265ac)
- Unit tests ([#99](https://github.com/reearth/reearth-backend/pull/99)) [`0d112c`](https://github.com/reearth/reearth-backend/commit/0d112c)
- Pkg/property, pkg/layer, pkg/plugin ([#101](https://github.com/reearth/reearth-backend/pull/101)) [`17a463`](https://github.com/reearth/reearth-backend/commit/17a463)

## 0.3.0 - 2022-01-11

### reearth-web

#### 🚀 Features

- Enhance terrain feature (type selection, exaggeration) ([#138](https://github.com/reearth/reearth-web/pull/138)) [`dae137`](https://github.com/reearth/reearth-web/commit/dae137)
- Clustering layers ([#143](https://github.com/reearth/reearth-web/pull/143)) [`3439cc`](https://github.com/reearth/reearth-web/commit/3439cc)
- Camera limiter ([#142](https://github.com/reearth/reearth-web/pull/142)) [`dec1dd`](https://github.com/reearth/reearth-web/commit/dec1dd)
- Tagging of layers ([#144](https://github.com/reearth/reearth-web/pull/144)) [`4d0a40`](https://github.com/reearth/reearth-web/commit/4d0a40)

#### 🔧 Bug Fixes

- Indicator is not displayed on selecting of clustered layer ([#146](https://github.com/reearth/reearth-web/pull/146)) [`e41f67`](https://github.com/reearth/reearth-web/commit/e41f67)
- Use data URL for marker images [`576ea4`](https://github.com/reearth/reearth-web/commit/576ea4)
- Layer clusters do not updated correctly [`ec74f6`](https://github.com/reearth/reearth-web/commit/ec74f6)
- Position label in front of billboard ([#147](https://github.com/reearth/reearth-web/pull/147)) [`81c533`](https://github.com/reearth/reearth-web/commit/81c533)
- Public pages do not work due to clustering feature [`48d8b3`](https://github.com/reearth/reearth-web/commit/48d8b3)
- Photooverlay transition does not work in Android ([#154](https://github.com/reearth/reearth-web/pull/154)) [`decbfe`](https://github.com/reearth/reearth-web/commit/decbfe)

#### 🎨 Styling

- Fix the height of the header [`9d6acc`](https://github.com/reearth/reearth-web/commit/9d6acc)

#### Miscellaneous Tasks

- Upgrade dependencies ([#134](https://github.com/reearth/reearth-web/pull/134)) [`740821`](https://github.com/reearth/reearth-web/commit/740821)
- Update dependency cesium to ^1.88.0 ([#139](https://github.com/reearth/reearth-web/pull/139)) [`7afdfb`](https://github.com/reearth/reearth-web/commit/7afdfb)
- Fix webpack dev server config [`8d06fa`](https://github.com/reearth/reearth-web/commit/8d06fa)
- Update dependency cesium to ^1.89.0 ([#156](https://github.com/reearth/reearth-web/pull/156)) [`d436ce`](https://github.com/reearth/reearth-web/commit/d436ce)

### reearth-backend

#### 🚀 Features

- Clusters for scenes ([#75](https://github.com/reearth/reearth-backend/pull/75)) [`3512c0`](https://github.com/reearth/reearth-backend/commit/3512c0)
- Add fields of scene property for terrain [`8693b4`](https://github.com/reearth/reearth-backend/commit/8693b4)
- Camera limiter  ([#87](https://github.com/reearth/reearth-backend/pull/87)) [`63c582`](https://github.com/reearth/reearth-backend/commit/63c582)

#### 🔧 Bug Fixes

- Terrain fields of scene property [`5e3d25`](https://github.com/reearth/reearth-backend/commit/5e3d25)
- Numbers are not decoded from gql to value [`2ddbc8`](https://github.com/reearth/reearth-backend/commit/2ddbc8)
- Layers have their own tags separate from the scene ([#90](https://github.com/reearth/reearth-backend/pull/90)) [`c4fb9a`](https://github.com/reearth/reearth-backend/commit/c4fb9a)
- Return property with clusters data ([#89](https://github.com/reearth/reearth-backend/pull/89)) [`1b99c6`](https://github.com/reearth/reearth-backend/commit/1b99c6)
- Cast values, rename value.OptionalValue ([#93](https://github.com/reearth/reearth-backend/pull/93)) [`ba4b18`](https://github.com/reearth/reearth-backend/commit/ba4b18)
- Synchronize mongo migration ([#94](https://github.com/reearth/reearth-backend/pull/94)) [`db4cea`](https://github.com/reearth/reearth-backend/commit/db4cea)

#### 📖 Documentation

- Add pkg.go.dev badge to readme [`91f9b3`](https://github.com/reearth/reearth-backend/commit/91f9b3)

#### ✨ Refactor

- Make property.Value and dataset.Value independent in pkg/value ([#77](https://github.com/reearth/reearth-backend/pull/77)) [`73143b`](https://github.com/reearth/reearth-backend/commit/73143b)

#### Miscellaneous Tasks

- Fix plugin manifest JSON schema [`2b57b1`](https://github.com/reearth/reearth-backend/commit/2b57b1)


## 0.2.0 - 2021-11-18

#### reearth-web

#### 🚀 Features

- Widget align system for mobile ([#115](https://github.com/reearth/reearth-web/pull/115)) [`afa4ba`](https://github.com/reearth/reearth-web/commit/afa4ba)
- Support dataset schema preview and create layer group from selected primitive type ([#74](https://github.com/reearth/reearth-web/pull/74)) [`769b86`](https://github.com/reearth/reearth-web/commit/769b86)

#### 🔧 Bug Fixes

- Markdown background color is not transparent ([#123](https://github.com/reearth/reearth-web/pull/123)) [`f16706`](https://github.com/reearth/reearth-web/commit/f16706)
- Layers would not be marshalled correctly ([#126](https://github.com/reearth/reearth-web/pull/126)) [`886302`](https://github.com/reearth/reearth-web/commit/886302)
- Widget align system issues ([#124](https://github.com/reearth/reearth-web/pull/124)) [`3bc9fa`](https://github.com/reearth/reearth-web/commit/3bc9fa)
- Project setting page does not display correctly after creating a new project ([#127](https://github.com/reearth/reearth-web/pull/127)) [`c120dc`](https://github.com/reearth/reearth-web/commit/c120dc)
- Dataset info pane shows its property though after selected dataset schema is deleted ([#131](https://github.com/reearth/reearth-web/pull/131)) [`2307d8`](https://github.com/reearth/reearth-web/commit/2307d8)

#### Miscellaneous Tasks

- Disable storybook workflow for release commit [`80f4d2`](https://github.com/reearth/reearth-web/commit/80f4d2)
- Change semantic commit type of renovate PRs, omit ci skip in changelog [`4a3e9e`](https://github.com/reearth/reearth-web/commit/4a3e9e)
- Follow backend GraphQL schema update ([#120](https://github.com/reearth/reearth-web/pull/120)) [`aeee1f`](https://github.com/reearth/reearth-web/commit/aeee1f)
- Load local reearth-config.json for debugging ([#119](https://github.com/reearth/reearth-web/pull/119)) [`6115ee`](https://github.com/reearth/reearth-web/commit/6115ee)
- Update dependency cesium to ^1.87.0 ([#118](https://github.com/reearth/reearth-web/pull/118)) [`7c65d0`](https://github.com/reearth/reearth-web/commit/7c65d0)
- Update dependency cesium to ^1.87.1 ([#128](https://github.com/reearth/reearth-web/pull/128)) [`a63aa7`](https://github.com/reearth/reearth-web/commit/a63aa7)
- Update codecov.yml to add ignored files [`b72f17`](https://github.com/reearth/reearth-web/commit/b72f17)


### reearth-backend

#### 🚀 Features

- Support opentelemetry ([#68](https://github.com/reearth/reearth-backend/pull/68)) [`25c581`](https://github.com/reearth/reearth-backend/commit/25c581)

#### 🔧 Bug Fixes

- Add an index to mongo project collection to prevent creating projects whose alias is duplicated [`10f745`](https://github.com/reearth/reearth-backend/commit/10f745)
- Check project alias duplication on project update [`443f2c`](https://github.com/reearth/reearth-backend/commit/443f2c)

#### ✨ Refactor

- Add PropertySchemaGroupID to pkg/id ([#70](https://github.com/reearth/reearth-backend/pull/70)) [`9ece9e`](https://github.com/reearth/reearth-backend/commit/9ece9e)

#### Miscellaneous Tasks

- Fix typo in github actions [`4a9dc5`](https://github.com/reearth/reearth-backend/commit/4a9dc5)
- Clean up unused code [`b5b01b`](https://github.com/reearth/reearth-backend/commit/b5b01b)
- Update codecov.yml to add ignored files [`d54309`](https://github.com/reearth/reearth-backend/commit/d54309)
- Ignore generated files in codecov [`9d3822`](https://github.com/reearth/reearth-backend/commit/9d3822)
- Upgrade dependencies [`215947`](https://github.com/reearth/reearth-backend/commit/215947)


## 0.1.0 - 2021-11-02

### reearth-web

#### 🚀 Features

- Support Auth0 audience ([#2](https://github.com/reearth/reearth-web/pull/2)) [`0ad0f6`](https://github.com/reearth/reearth-web/commit/0ad0f6)
- Asset modal redesign ([#1](https://github.com/reearth/reearth-web/pull/1)) [`f71117`](https://github.com/reearth/reearth-web/commit/f71117)
- Basic auth for projects ([#3](https://github.com/reearth/reearth-web/pull/3)) [`372c4e`](https://github.com/reearth/reearth-web/commit/372c4e)
- Google analytics ([#6](https://github.com/reearth/reearth-web/pull/6)) [`01aadf`](https://github.com/reearth/reearth-web/commit/01aadf)
- Refine setting page ([#19](https://github.com/reearth/reearth-web/pull/19)) [`d06ee7`](https://github.com/reearth/reearth-web/commit/d06ee7)
- Add  delete assets confirm modal and fix bugs ([#25](https://github.com/reearth/reearth-web/pull/25)) [`0310f5`](https://github.com/reearth/reearth-web/commit/0310f5)
- Update link system and UI ([#12](https://github.com/reearth/reearth-web/pull/12)) [`51de77`](https://github.com/reearth/reearth-web/commit/51de77)
- Import google sheet dataset ([#14](https://github.com/reearth/reearth-web/pull/14)) [`21b167`](https://github.com/reearth/reearth-web/commit/21b167)
- Support multi-line infobox titles ([#40](https://github.com/reearth/reearth-web/pull/40)) [`4cddcc`](https://github.com/reearth/reearth-web/commit/4cddcc)
- Public settings page ([#32](https://github.com/reearth/reearth-web/pull/32)) [`ebfd41`](https://github.com/reearth/reearth-web/commit/ebfd41)
- Refine readme ([#51](https://github.com/reearth/reearth-web/pull/51)) [`41ddb3`](https://github.com/reearth/reearth-web/commit/41ddb3)
- Add light theme ([#52](https://github.com/reearth/reearth-web/pull/52)) [`26159b`](https://github.com/reearth/reearth-web/commit/26159b)
- Add a short discription for light theme ([#56](https://github.com/reearth/reearth-web/pull/56)) [`8b092d`](https://github.com/reearth/reearth-web/commit/8b092d)
- Plugins settings page, install/uninstall plugins ([#22](https://github.com/reearth/reearth-web/pull/22)) [`018674`](https://github.com/reearth/reearth-web/commit/018674)
- Plugin system, refactoring visualizer ([#50](https://github.com/reearth/reearth-web/pull/50)) [`172939`](https://github.com/reearth/reearth-web/commit/172939)
- 3D tileset, model, rectangle primitive, more properties for marker and scene ([#63](https://github.com/reearth/reearth-web/pull/63)) [`a88600`](https://github.com/reearth/reearth-web/commit/a88600)
- Graphiql page ([#70](https://github.com/reearth/reearth-web/pull/70)) [`aa5d10`](https://github.com/reearth/reearth-web/commit/aa5d10)
- Enable to set theme for the scene ([#67](https://github.com/reearth/reearth-web/pull/67)) [`58e670`](https://github.com/reearth/reearth-web/commit/58e670)
- Notification system update ([#73](https://github.com/reearth/reearth-web/pull/73)) [`92cdbb`](https://github.com/reearth/reearth-web/commit/92cdbb)
- Widget align system ([#61](https://github.com/reearth/reearth-web/pull/61)) [`ed2940`](https://github.com/reearth/reearth-web/commit/ed2940)
- Plugin system beta ([#87](https://github.com/reearth/reearth-web/pull/87)) [`d76f1c`](https://github.com/reearth/reearth-web/commit/d76f1c)
- Enhance extended field of widget in plugin API ([#90](https://github.com/reearth/reearth-web/pull/90)) [`06cb14`](https://github.com/reearth/reearth-web/commit/06cb14)
- Add overrideProperty to plugin layers API ([#92](https://github.com/reearth/reearth-web/pull/92)) [`563f88`](https://github.com/reearth/reearth-web/commit/563f88)
- Add a fallback icon for extensions that do not have an icon ([#98](https://github.com/reearth/reearth-web/pull/98)) [`50de1f`](https://github.com/reearth/reearth-web/commit/50de1f)
- Add password validation ([#86](https://github.com/reearth/reearth-web/pull/86)) [`2017aa`](https://github.com/reearth/reearth-web/commit/2017aa)
- Draggable layer ([#58](https://github.com/reearth/reearth-web/pull/58)) [`25a217`](https://github.com/reearth/reearth-web/commit/25a217)
- Multi widgets ([#99](https://github.com/reearth/reearth-web/pull/99)) [`bea1a3`](https://github.com/reearth/reearth-web/commit/bea1a3)
- Front-end for new authentication system ([#102](https://github.com/reearth/reearth-web/pull/102)) [`964d92`](https://github.com/reearth/reearth-web/commit/964d92)
- Add layers.overriddenProperties, find, findAll, walk ([#110](https://github.com/reearth/reearth-web/pull/110)) [`ebe131`](https://github.com/reearth/reearth-web/commit/ebe131)

#### 🔧 Bug Fixes

- Reorganize config [`f2e947`](https://github.com/reearth/reearth-web/commit/f2e947)
- Update gql schema [`0905b6`](https://github.com/reearth/reearth-web/commit/0905b6)
- Update dependency cesium to ^1.82.1 ([#4](https://github.com/reearth/reearth-web/pull/4)) [`0627bf`](https://github.com/reearth/reearth-web/commit/0627bf)
- Google analytics ([#7](https://github.com/reearth/reearth-web/pull/7)) [`7505ca`](https://github.com/reearth/reearth-web/commit/7505ca)
- Sprint15 bugs ([#8](https://github.com/reearth/reearth-web/pull/8)) [`e2fe0a`](https://github.com/reearth/reearth-web/commit/e2fe0a)
- Google analytics typo ([#9](https://github.com/reearth/reearth-web/pull/9)) [`943b5e`](https://github.com/reearth/reearth-web/commit/943b5e)
- Ga-typo2 ([#10](https://github.com/reearth/reearth-web/pull/10)) [`b498de`](https://github.com/reearth/reearth-web/commit/b498de)
- Force logout when me query returns null ([#15](https://github.com/reearth/reearth-web/pull/15)) [`339d61`](https://github.com/reearth/reearth-web/commit/339d61)
- Infinit logout loop ([#17](https://github.com/reearth/reearth-web/pull/17)) [`0d510f`](https://github.com/reearth/reearth-web/commit/0d510f)
- Change data.json path [`38a69a`](https://github.com/reearth/reearth-web/commit/38a69a)
- Menu button width ([#21](https://github.com/reearth/reearth-web/pull/21)) [`d08eba`](https://github.com/reearth/reearth-web/commit/d08eba)
- Menu widget bugs ([#37](https://github.com/reearth/reearth-web/pull/37)) [`5d5483`](https://github.com/reearth/reearth-web/commit/5d5483)
- Marker label position is oposite to actual display ([#39](https://github.com/reearth/reearth-web/pull/39)) [`38de46`](https://github.com/reearth/reearth-web/commit/38de46)
- Disable default cesium mouse event ([#42](https://github.com/reearth/reearth-web/pull/42)) [`129ae3`](https://github.com/reearth/reearth-web/commit/129ae3)
- Show layers in storytelling without names ([#45](https://github.com/reearth/reearth-web/pull/45)) [`00ae3c`](https://github.com/reearth/reearth-web/commit/00ae3c)
- Infobox colors ([#47](https://github.com/reearth/reearth-web/pull/47)) [`2a6a36`](https://github.com/reearth/reearth-web/commit/2a6a36)
- Project public image ([#48](https://github.com/reearth/reearth-web/pull/48)) [`91b5ee`](https://github.com/reearth/reearth-web/commit/91b5ee)
- Auth0 redirect uri [`8336a3`](https://github.com/reearth/reearth-web/commit/8336a3)
- Storybook ([#54](https://github.com/reearth/reearth-web/pull/54)) [`fde0c0`](https://github.com/reearth/reearth-web/commit/fde0c0)
- Published data url [`e3d5b0`](https://github.com/reearth/reearth-web/commit/e3d5b0)
- Icon background ([#64](https://github.com/reearth/reearth-web/pull/64)) [`9c69a4`](https://github.com/reearth/reearth-web/commit/9c69a4)
- Prevent extra render, cannot rename layers, cannot display infobox on dataset layers ([#65](https://github.com/reearth/reearth-web/pull/65)) [`e3d618`](https://github.com/reearth/reearth-web/commit/e3d618)
- Remove visibility icon from layer actions [`0ad8aa`](https://github.com/reearth/reearth-web/commit/0ad8aa)
- Default published url, rename layer when focus is removed from text box [`f9accc`](https://github.com/reearth/reearth-web/commit/f9accc)
- Storybook error ([#75](https://github.com/reearth/reearth-web/pull/75)) [`f27f9b`](https://github.com/reearth/reearth-web/commit/f27f9b)
- Showing members section for personal workspace ([#85](https://github.com/reearth/reearth-web/pull/85)) [`8e78f9`](https://github.com/reearth/reearth-web/commit/8e78f9)
- Widget bugs, language ([#89](https://github.com/reearth/reearth-web/pull/89)) [`9de9df`](https://github.com/reearth/reearth-web/commit/9de9df)
- Update dependency cesium to ^1.86.0 ([#93](https://github.com/reearth/reearth-web/pull/93)) [`7ca298`](https://github.com/reearth/reearth-web/commit/7ca298)
- Show properties of 3D tile features on infobox ([#95](https://github.com/reearth/reearth-web/pull/95)) [`a9cc23`](https://github.com/reearth/reearth-web/commit/a9cc23)
- Navigator.language should be used as fallback lang ([#91](https://github.com/reearth/reearth-web/pull/91)) [`15df16`](https://github.com/reearth/reearth-web/commit/15df16)
- Camera property panel bugs ([#96](https://github.com/reearth/reearth-web/pull/96)) [`2c3eaa`](https://github.com/reearth/reearth-web/commit/2c3eaa)
- Camera flight bugs ([#97](https://github.com/reearth/reearth-web/pull/97)) [`b4f1ae`](https://github.com/reearth/reearth-web/commit/b4f1ae)
- Storytelling image crop does not work [`9c23b3`](https://github.com/reearth/reearth-web/commit/9c23b3)
- Export pane is not displayed [`58ceda`](https://github.com/reearth/reearth-web/commit/58ceda)
- 1st bug hunt of october ([#100](https://github.com/reearth/reearth-web/pull/100)) [`1b9032`](https://github.com/reearth/reearth-web/commit/1b9032)
- Layers disappearing when in nested folders ([#101](https://github.com/reearth/reearth-web/pull/101)) [`778395`](https://github.com/reearth/reearth-web/commit/778395)
- Update dependency cesium to ^1.86.1 ([#103](https://github.com/reearth/reearth-web/pull/103)) [`385582`](https://github.com/reearth/reearth-web/commit/385582)
- Bug bounty #2 ([#105](https://github.com/reearth/reearth-web/pull/105)) [`da4815`](https://github.com/reearth/reearth-web/commit/da4815)
- Button widget ([#111](https://github.com/reearth/reearth-web/pull/111)) [`b93485`](https://github.com/reearth/reearth-web/commit/b93485)
- Create team redirect + translations update ([#112](https://github.com/reearth/reearth-web/pull/112)) [`bafcfd`](https://github.com/reearth/reearth-web/commit/bafcfd)
- 3d tile styles not updating sometimes ([#109](https://github.com/reearth/reearth-web/pull/109)) [`1e92b8`](https://github.com/reearth/reearth-web/commit/1e92b8)
- Layers.overrideProperty property merging and rerendering ([#108](https://github.com/reearth/reearth-web/pull/108)) [`e5c255`](https://github.com/reearth/reearth-web/commit/e5c255)
- Password policy conversion in config ([#113](https://github.com/reearth/reearth-web/pull/113)) [`5d57c4`](https://github.com/reearth/reearth-web/commit/5d57c4)
- Password validation, add autofocus ([#117](https://github.com/reearth/reearth-web/pull/117)) [`348454`](https://github.com/reearth/reearth-web/commit/348454)
- Password verification, add better feedback [`bd1725`](https://github.com/reearth/reearth-web/commit/bd1725)

#### ✨ Refactor

- Use jotai instead of redux ([#68](https://github.com/reearth/reearth-web/pull/68)) [`ea980c`](https://github.com/reearth/reearth-web/commit/ea980c)
- Replace deprecated gql fields, pass widgetId to widget mutations ([#72](https://github.com/reearth/reearth-web/pull/72)) [`f36c86`](https://github.com/reearth/reearth-web/commit/f36c86)

#### 🎨 Styling

- Refine font ([#49](https://github.com/reearth/reearth-web/pull/49)) [`8b3755`](https://github.com/reearth/reearth-web/commit/8b3755)
- Refine color vo.1 ([#59](https://github.com/reearth/reearth-web/pull/59)) [`ab7bce`](https://github.com/reearth/reearth-web/commit/ab7bce)

#### 🧪 Testing

- Fix e2e test [`3bcd2d`](https://github.com/reearth/reearth-web/commit/3bcd2d)
- Fix e2e test [`b3e512`](https://github.com/reearth/reearth-web/commit/b3e512)
- Fix e2e test [`277f4e`](https://github.com/reearth/reearth-web/commit/277f4e)
- Fix e2e test [`396f71`](https://github.com/reearth/reearth-web/commit/396f71)
- Fix e2e test [`a8bd0c`](https://github.com/reearth/reearth-web/commit/a8bd0c)
- Fix e2e test [`fd7cf5`](https://github.com/reearth/reearth-web/commit/fd7cf5)
- Fix e2e test [`8c300b`](https://github.com/reearth/reearth-web/commit/8c300b)
- Fix e2e test [`ea5050`](https://github.com/reearth/reearth-web/commit/ea5050)
- Fix e2e test [`866c8c`](https://github.com/reearth/reearth-web/commit/866c8c)
- Support display name in e2e test [`0edf58`](https://github.com/reearth/reearth-web/commit/0edf58)

#### Miscellaneous Tasks

- Update workflows, set up nightly release [`0ea0ff`](https://github.com/reearth/reearth-web/commit/0ea0ff)
- Fix nightly release workflow [`d7d1d3`](https://github.com/reearth/reearth-web/commit/d7d1d3)
- Fix config [`7a6ed3`](https://github.com/reearth/reearth-web/commit/7a6ed3)
- Set up cd workflows [`a6f0f5`](https://github.com/reearth/reearth-web/commit/a6f0f5)
- Fix workflows [`97ecf8`](https://github.com/reearth/reearth-web/commit/97ecf8)
- Fix workflows [`a4d451`](https://github.com/reearth/reearth-web/commit/a4d451)
- Fix workflows [`d77b53`](https://github.com/reearth/reearth-web/commit/d77b53)
- Remove unused deps [`81d0eb`](https://github.com/reearth/reearth-web/commit/81d0eb)
- Update cesium [`414b37`](https://github.com/reearth/reearth-web/commit/414b37)
- Update renovate config [`b36740`](https://github.com/reearth/reearth-web/commit/b36740)
- Use .env instead of .env.local [`0b8720`](https://github.com/reearth/reearth-web/commit/0b8720)
- Add storybook workflow [`c624bd`](https://github.com/reearth/reearth-web/commit/c624bd)
- Set up sentry ([#18](https://github.com/reearth/reearth-web/pull/18)) [`8a2d38`](https://github.com/reearth/reearth-web/commit/8a2d38)
- Testable published page ([#43](https://github.com/reearth/reearth-web/pull/43)) [`90c37d`](https://github.com/reearth/reearth-web/commit/90c37d)
- Update netlify.toml [`230e12`](https://github.com/reearth/reearth-web/commit/230e12)
- Add gql sclar types [`09fb76`](https://github.com/reearth/reearth-web/commit/09fb76)
- Update cesium and resium ([#79](https://github.com/reearth/reearth-web/pull/79)) [`c41601`](https://github.com/reearth/reearth-web/commit/c41601)
- Update eslint, enable eslint-plugin-import, perform formatting ([#82](https://github.com/reearth/reearth-web/pull/82)) [`117bab`](https://github.com/reearth/reearth-web/commit/117bab)
- Upgrade dependencies [`4924f9`](https://github.com/reearth/reearth-web/commit/4924f9)
- Fix cypress and unit test [`97f74e`](https://github.com/reearth/reearth-web/commit/97f74e)
- Support for dotenv switching ([#106](https://github.com/reearth/reearth-web/pull/106)) [`cd1974`](https://github.com/reearth/reearth-web/commit/cd1974)
- Upgrade modules oct ([#107](https://github.com/reearth/reearth-web/pull/107)) [`24c145`](https://github.com/reearth/reearth-web/commit/24c145)
- Upgrade react-align ([#116](https://github.com/reearth/reearth-web/pull/116)) [`7f4b98`](https://github.com/reearth/reearth-web/commit/7f4b98)
- Add github workflows to release [`331afb`](https://github.com/reearth/reearth-web/commit/331afb)
- Update translations + format ([#114](https://github.com/reearth/reearth-web/pull/114)) [`7f191e`](https://github.com/reearth/reearth-web/commit/7f191e)
- Lock file maintenance ([#66](https://github.com/reearth/reearth-web/pull/66)) [`6d2a2d`](https://github.com/reearth/reearth-web/commit/6d2a2d)
- Fix slack notifications in workflow [skip ci] [`b4fa35`](https://github.com/reearth/reearth-web/commit/b4fa35)
- Fix sed in release workflow [skip ci] [`f3cd74`](https://github.com/reearth/reearth-web/commit/f3cd74)


### reearth-backend

#### 🚀 Features

- Support Auth0 audience ([#3](https://github.com/reearth/reearth-backend/pull/3)) [`c3758e`](https://github.com/reearth/reearth-backend/commit/c3758e)
- Basic auth for projects ([#6](https://github.com/reearth/reearth-backend/pull/6)) [`5db065`](https://github.com/reearth/reearth-backend/commit/5db065)
- Google analytics for scene ([#10](https://github.com/reearth/reearth-backend/pull/10)) [`b44249`](https://github.com/reearth/reearth-backend/commit/b44249)
- Create installable plugins ([#1](https://github.com/reearth/reearth-backend/pull/1)) [`5b7a5f`](https://github.com/reearth/reearth-backend/commit/5b7a5f)
- Add thumbnail, author fields on plugin metadata query ([#15](https://github.com/reearth/reearth-backend/pull/15)) [`888fe0`](https://github.com/reearth/reearth-backend/commit/888fe0)
- Published page api ([#11](https://github.com/reearth/reearth-backend/pull/11)) [`aebac3`](https://github.com/reearth/reearth-backend/commit/aebac3)
- Import dataset from google sheets ([#16](https://github.com/reearth/reearth-backend/pull/16)) [`2ef7ef`](https://github.com/reearth/reearth-backend/commit/2ef7ef)
- Add scenePlugin resolver to layers ([#20](https://github.com/reearth/reearth-backend/pull/20)) [`5213f3`](https://github.com/reearth/reearth-backend/commit/5213f3)
- Marker label position [`bb9e4c`](https://github.com/reearth/reearth-backend/commit/bb9e4c)
- Refine dataset import ([#26](https://github.com/reearth/reearth-backend/pull/26)) [`5dd3db`](https://github.com/reearth/reearth-backend/commit/5dd3db)
- Plugin upload and deletion ([#33](https://github.com/reearth/reearth-backend/pull/33)) [`8742db`](https://github.com/reearth/reearth-backend/commit/8742db)
- New primitives, new properties on primitives [`108711`](https://github.com/reearth/reearth-backend/commit/108711)
- Set scene theme ([#35](https://github.com/reearth/reearth-backend/pull/35)) [`2e4f52`](https://github.com/reearth/reearth-backend/commit/2e4f52)
- Widget align system ([#19](https://github.com/reearth/reearth-backend/pull/19)) [`94611f`](https://github.com/reearth/reearth-backend/commit/94611f)
- Tag system ([#67](https://github.com/reearth/reearth-backend/pull/67)) [`163fcf`](https://github.com/reearth/reearth-backend/commit/163fcf)

#### 🔧 Bug Fixes

- Add mutex for each memory repo ([#2](https://github.com/reearth/reearth-backend/pull/2)) [`f4c3b0`](https://github.com/reearth/reearth-backend/commit/f4c3b0)
- Auth0 audience in reearth_config.json [`72e3ed`](https://github.com/reearth/reearth-backend/commit/72e3ed)
- Auth0 domain and multiple auds [`835a02`](https://github.com/reearth/reearth-backend/commit/835a02)
- Signing up and deleting user [`f17b9d`](https://github.com/reearth/reearth-backend/commit/f17b9d)
- Deleting user [`e9b8c9`](https://github.com/reearth/reearth-backend/commit/e9b8c9)
- Sign up and update user [`e5ab87`](https://github.com/reearth/reearth-backend/commit/e5ab87)
- Make gql mutation payloads optional [`9b1c4a`](https://github.com/reearth/reearth-backend/commit/9b1c4a)
- Auth0 [`6a27c6`](https://github.com/reearth/reearth-backend/commit/6a27c6)
- Errors are be overwriten by tx [`2d08c5`](https://github.com/reearth/reearth-backend/commit/2d08c5)
- Deleting user [`f531bd`](https://github.com/reearth/reearth-backend/commit/f531bd)
- Always enable dev mode in debug [`0815d3`](https://github.com/reearth/reearth-backend/commit/0815d3)
- User deletion [`a5eeae`](https://github.com/reearth/reearth-backend/commit/a5eeae)
- Invisible layer issue in published project ([#7](https://github.com/reearth/reearth-backend/pull/7)) [`06cd44`](https://github.com/reearth/reearth-backend/commit/06cd44)
- Dataset link merge bug #378 ([#18](https://github.com/reearth/reearth-backend/pull/18)) [`25da0d`](https://github.com/reearth/reearth-backend/commit/25da0d)
- Ogp image for published page ([#17](https://github.com/reearth/reearth-backend/pull/17)) [`dcb4b0`](https://github.com/reearth/reearth-backend/commit/dcb4b0)
- Change default value of marker label position [`a2059e`](https://github.com/reearth/reearth-backend/commit/a2059e)
- Import dataset from google sheet bug ([#23](https://github.com/reearth/reearth-backend/pull/23)) [`077558`](https://github.com/reearth/reearth-backend/commit/077558)
- Public api param [`846957`](https://github.com/reearth/reearth-backend/commit/846957)
- Replace strings.Split() with strings.field() ([#25](https://github.com/reearth/reearth-backend/pull/25)) [`ba7d16`](https://github.com/reearth/reearth-backend/commit/ba7d16)
- Project public image type [`e82b54`](https://github.com/reearth/reearth-backend/commit/e82b54)
- Published API ([#27](https://github.com/reearth/reearth-backend/pull/27)) [`8ad1f8`](https://github.com/reearth/reearth-backend/commit/8ad1f8)
- Plugin manifest parser bugs ([#32](https://github.com/reearth/reearth-backend/pull/32)) [`78ac13`](https://github.com/reearth/reearth-backend/commit/78ac13)
- Dataset layers are not exported correctly ([#36](https://github.com/reearth/reearth-backend/pull/36)) [`0b8c00`](https://github.com/reearth/reearth-backend/commit/0b8c00)
- Hide parent infobox fields when child infobox is not nil ([#37](https://github.com/reearth/reearth-backend/pull/37)) [`d8c8cd`](https://github.com/reearth/reearth-backend/commit/d8c8cd)
- Mongo.PropertySchema.FindByIDs, propertySchemaID.Equal [`be00da`](https://github.com/reearth/reearth-backend/commit/be00da)
- Gql propertySchemaGroup.translatedTitle resolver [`a4770e`](https://github.com/reearth/reearth-backend/commit/a4770e)
- Use PropertySchemaID.Equal [`8a6459`](https://github.com/reearth/reearth-backend/commit/8a6459)
- Use PropertySchemaID.Equal [`1c3cf1`](https://github.com/reearth/reearth-backend/commit/1c3cf1)
- Tweak field names of model primitive [`080ab9`](https://github.com/reearth/reearth-backend/commit/080ab9)
- Layer importing bug ([#41](https://github.com/reearth/reearth-backend/pull/41)) [`02b17f`](https://github.com/reearth/reearth-backend/commit/02b17f)
- Skip nil geometries ([#42](https://github.com/reearth/reearth-backend/pull/42)) [`90c327`](https://github.com/reearth/reearth-backend/commit/90c327)
- Validate widget extended when moved [`a7daf7`](https://github.com/reearth/reearth-backend/commit/a7daf7)
- Widget extended validation [`98db7e`](https://github.com/reearth/reearth-backend/commit/98db7e)
- Nil error in mongodoc plugin [`d236be`](https://github.com/reearth/reearth-backend/commit/d236be)
- Add widget to default location [`eb1db4`](https://github.com/reearth/reearth-backend/commit/eb1db4)
- Invalid extension data from GraphQL, plugin manifest schema improvement, more friendly error from manifest parser ([#56](https://github.com/reearth/reearth-backend/pull/56)) [`92d137`](https://github.com/reearth/reearth-backend/commit/92d137)
- Translated fields in plugin gql [`0a658a`](https://github.com/reearth/reearth-backend/commit/0a658a)
- Fallback widgetLocation [`579b7a`](https://github.com/reearth/reearth-backend/commit/579b7a)

#### 📖 Documentation

- Refine readme ([#28](https://github.com/reearth/reearth-backend/pull/28)) [`a9d209`](https://github.com/reearth/reearth-backend/commit/a9d209)
- Add badges to readme [skip ci] [`cc63cd`](https://github.com/reearth/reearth-backend/commit/cc63cd)

#### ✨ Refactor

- Remove unused code [`37b2c2`](https://github.com/reearth/reearth-backend/commit/37b2c2)
- Pkg/error ([#31](https://github.com/reearth/reearth-backend/pull/31)) [`a3f8b6`](https://github.com/reearth/reearth-backend/commit/a3f8b6)
- Graphql adapter ([#40](https://github.com/reearth/reearth-backend/pull/40)) [`2a1d4f`](https://github.com/reearth/reearth-backend/commit/2a1d4f)
- Reorganize graphql schema ([#43](https://github.com/reearth/reearth-backend/pull/43)) [`d3360b`](https://github.com/reearth/reearth-backend/commit/d3360b)

#### 🧪 Testing

- Pkg/shp ([#5](https://github.com/reearth/reearth-backend/pull/5)) [`72ed8e`](https://github.com/reearth/reearth-backend/commit/72ed8e)
- Pkg/id ([#4](https://github.com/reearth/reearth-backend/pull/4)) [`c31bdb`](https://github.com/reearth/reearth-backend/commit/c31bdb)

#### Miscellaneous Tasks

- Enable nightly release workflow [`16c037`](https://github.com/reearth/reearth-backend/commit/16c037)
- Set up workflows [`819639`](https://github.com/reearth/reearth-backend/commit/819639)
- Fix workflows [`c022a4`](https://github.com/reearth/reearth-backend/commit/c022a4)
- Print config [`0125aa`](https://github.com/reearth/reearth-backend/commit/0125aa)
- Load .env instead of .env.local [`487a73`](https://github.com/reearth/reearth-backend/commit/487a73)
- Add godoc workflow [`9629dd`](https://github.com/reearth/reearth-backend/commit/9629dd)
- Fix godoc workflow [`cc45b5`](https://github.com/reearth/reearth-backend/commit/cc45b5)
- Fix godoc workflow [`0db163`](https://github.com/reearth/reearth-backend/commit/0db163)
- Fix godoc workflow [`9b78fc`](https://github.com/reearth/reearth-backend/commit/9b78fc)
- Fix godoc workflow [`f1e5a7`](https://github.com/reearth/reearth-backend/commit/f1e5a7)
- Fix godoc workflow [`f7866c`](https://github.com/reearth/reearth-backend/commit/f7866c)
- Fix godoc workflow [`5bc089`](https://github.com/reearth/reearth-backend/commit/5bc089)
- Fix godoc workflow [`5f808b`](https://github.com/reearth/reearth-backend/commit/5f808b)
- Fix godoc workflow [`9f8e11`](https://github.com/reearth/reearth-backend/commit/9f8e11)
- Fix godoc workflow [`150550`](https://github.com/reearth/reearth-backend/commit/150550)
- Use go:embed ([#24](https://github.com/reearth/reearth-backend/pull/24)) [`f7866e`](https://github.com/reearth/reearth-backend/commit/f7866e)
- Add internal error log [`41c377`](https://github.com/reearth/reearth-backend/commit/41c377)
- Support multiple platform docker image [`3651e2`](https://github.com/reearth/reearth-backend/commit/3651e2)
- Stop using upx as it doesn't work on arm64 [`3b5f93`](https://github.com/reearth/reearth-backend/commit/3b5f93)
- Update golang version and modules ([#51](https://github.com/reearth/reearth-backend/pull/51)) [`33f4c7`](https://github.com/reearth/reearth-backend/commit/33f4c7)
- Updating modules ([#62](https://github.com/reearth/reearth-backend/pull/62)) [`65ae32`](https://github.com/reearth/reearth-backend/commit/65ae32)
- Add github workflows to release [`fbcdef`](https://github.com/reearth/reearth-backend/commit/fbcdef)
- Fix release workflow, fix build comment [skip ci] [`cfc79a`](https://github.com/reearth/reearth-backend/commit/cfc79a)
- Fix renaming file names in release workflow [`96f0b3`](https://github.com/reearth/reearth-backend/commit/96f0b3)
- Fix and refactor release workflow [skip ci] [`d5466b`](https://github.com/reearth/reearth-backend/commit/d5466b)