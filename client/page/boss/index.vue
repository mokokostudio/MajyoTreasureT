<template>
  <div class="boss-view">
    <!-- headbox -->
    <headers />

    <!-- weaponlist -->
    <div class="bossbox">
      <div class="bosslist">
        <div class="my-buff">
          <div v-for="(buff, index) in this.buffCards" :key="index" @click="onPopBox(6)">
            <img :src="buff.img" />
            <span class="text-stroke">{{buff.leftround}}</span>
          </div>
        </div>
        <div :class="bossChangeMode(this.bossChange)">
          <div class="teitem" :class="swipeMode(this.bossAni)">
            <div v-if="challenge && nftBossChallenge.equipSameNft && nftBossChallenge.upgradeOK" class="bospic" @click="onPopBox(7)">
              <img :src="this.thisBossData.picPath" />
            </div>
            <div v-else class="bospic-unable" @click="onPopBox(7)">
              <img :src="this.thisBossData.picShadowPath" />
              <div v-if="this.thisBossData.Class == 0" class="boss-tip">
                <div class="tip-title text-stroke">Conditions</div>
                <div class="tip-ct text-stroke">Defeat previous BOSS</div>
              </div>
              <div v-if="this.thisBossData.Class > 0" class="boss-tip">
                <div class="tip-title nft-boss text-stroke">Conditions</div>
                <div v-if="!nftBossChallenge.firstBoss" :class="challenge ? 'tip-ct pass text-stroke' : 'tip-ct text-stroke'">Defeat previous BOSS<em v-if="challenge" class="cond-ok"></em></div>
                <div :class="nftBossChallenge.upgradeOK ? 'tip-ct pass text-stroke' : 'tip-ct text-stroke'">Equipped NFT level<em class="cond-lv"></em>{{this.thisBossData.NFTEquipsLevel}}<em v-if="nftBossChallenge.upgradeOK" class="cond-ok"></em></div>
                <div :class="nftBossChallenge.equipSameNft ? 'tip-ct pass text-stroke' : 'tip-ct text-stroke'">
                  Equip an NFT
                  <img v-if="this.thisBossData.Class == 1" class="cond-nft" src="~@/assets/nfts/Garm6.png">
                  <img v-if="this.thisBossData.Class == 2" class="cond-nft" src="~@/assets/nfts/Norns6.png">
                  <img v-if="this.thisBossData.Class == 3" class="cond-nft" src="~@/assets/nfts/Tartarus6.png">
                  <img v-if="this.thisBossData.Class == 4" class="cond-nft" src="~@/assets/nfts/Parvati6.png">
                  <img v-if="this.thisBossData.Class == 5" class="cond-nft" src="~@/assets/nfts/Edda6.png">
                  <em v-if="nftBossChallenge.equipSameNft" class="cond-ok"></em>
                </div>
              </div>
            </div>
            <div class="boslove">
              <span class="text-stroke" @click="onPopBox(7)">{{this.thisBossData.nameStr}} {{ this.thisBossData.LevelShow }}</span>
              <img src="../../assets/info.png" @click="onPopBox(7)"/>
            </div>
            <div class="first-awards">
              <div class="rewards-title text-stroke">
                <span>Rewards</span>
              </div>
              <div class="rewards-content">
                <div class="rds-icon" v-for="item in this.firstRewardsData" :key="item.itemId">
                  <img v-if="showcheck" class="rewards-check-ok" src="../../assets/boss/yes.png" />
                  <img v-if="showcheck" class="rewards-check" src="../../assets/tipinfo.png" />
                  <img :src="item.itemPic" />
                  <div v-if="!showcheck" class="itemNumNoAdd text-stroke" v-html="item.num"></div>
                </div>
                <div class="rds-icon" v-for="item in this.thisBossData.Awards" :key="item.itemId">
                  <img :src="item.itemPic" />
                  <div class="itemNum text-stroke">{{item.min}}<span>+</span></div>
                </div>
                <div v-if="this.thisBossData.ManaAwardsFinal > 0" class="rds-icon">
                  <img src="~@/assets/ETHERicon.png" />
                  <div class="itemNumNoAdd text-stroke">{{this.thisBossData.ManaAwardsFinal}}</div>
                </div>
              </div>
            </div>
            <div class="tonvalue">
              <div class="text-stroke" v-for="item in this.thisBossData.Awards" :key="item.itemId">
                <span>Rewards value ≈ {{item.max * 0.5}} U</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Optional controls -->

        <div class="swiper-button-prev"  @click="buttonPrev()"><span></span></div>
        <div class="swiper-button-next"  @click="buttonNext()"><span></span></div>


        <div class="boicon">
          <div :class="ons === 1 ? 'itemsd ons' : 'itemsd'" @click="buttonOne(1)"><img alt="NFT"
              src="~@/assets/boss/boss-tab-1.png" /></div>
          <div :class="ons === 2 ? 'itemsd ons' : 'itemsd'" @click="buttonOne(2)"><img alt="NFT"
              src="~@/assets/boss/boss-tab-2.png" /></div>
          <div :class="ons === 3 ? 'itemsd ons' : 'itemsd'" @click="buttonOne(3)"><img alt="NFT"
              src="~@/assets/boss/boss-tab-3.png" /></div>
          <div :class="ons === 4 ? 'itemsd ons' : 'itemsd'" @click="buttonOne(4)"><img alt="NFT"
              src="~@/assets/boss/boss-tab-4.png" /></div>
          <div :class="ons === 5 ? 'itemsd ons' : 'itemsd'" @click="buttonOne(5)"><img alt="NFT"
              src="~@/assets/boss/boss-tab-5.png" /></div>
          <div :class="ons === 6 ? 'itemsd ons' : 'itemsd'" @click="buttonOne(6)"><img alt="NFT"
              src="~@/assets/boss/boss-tab-6.png" /></div>
        </div>

        <div v-if="challenge && nftBossChallenge.equipSameNft && nftBossChallenge.upgradeOK">
          <div class="costEnergy">
            <img src="~@/assets/items/92000001.png" />
            <div class="energyNum"><span class="text-stroke">X</span><span class="text-stroke">{{ this.thisBossData.EnergyCost }}</span></div>
          </div>
          <div class="stsrtbtn"><span @click="onStartBtn()"></span></div>
        </div>
        <div v-else>
          <div class="nostsrtbtn"><span></span></div>
        </div>
      </div>
    </div>

    <div class="iconls">
      <div class="sbx">
        <div class="pic baseitem" @click="onSbx(sbxData[0].equip_type)">
          <img v-if="sbxData[0].level < 3" alt="NFT" src="~@/assets/boss/ordinary1.png" />
          <img v-if="sbxData[0].level >= 3 && sbxData[0].level < 9" alt="NFT" src="~@/assets/boss/ordinary11.png" />
          <img v-if="sbxData[0].level >= 9" alt="NFT" src="~@/assets/boss/ordinary12.png" />
          <div class="star"><span class="text-stroke">{{ sbxData[0].star }}</span></div>
          <div class="lv"><span class="text-stroke">Lv.{{ sbxData[0].level }}</span></div>
          <div class="redpoint" v-if="redpoint.equipLv1 || redpoint.equipStar1"></div>
        </div>
        <div class="pic baseitem" @click="onSbx(sbxData[1].equip_type)">
          <img v-if="sbxData[1].level < 3" alt="NFT" src="~@/assets/boss/ordinary2.png" />
          <img v-if="sbxData[1].level >= 3 && sbxData[1].level < 9" alt="NFT" src="~@/assets/boss/ordinary21.png" />
          <img v-if="sbxData[1].level >= 9" alt="NFT" src="~@/assets/boss/ordinary22.png" />
          <div class="star"><span class="text-stroke">{{ sbxData[1].star }}</span></div>
          <div class="lv"><span class="text-stroke">Lv.{{ sbxData[1].level }}</span></div>
          <div class="redpoint" v-if="redpoint.equipLv2 || redpoint.equipStar2"></div>
        </div>
        <div class="pic baseitem" @click="onSbx(sbxData[2].equip_type)">
          <img v-if="sbxData[2].level < 3" alt="NFT" src="~@/assets/boss/ordinary3.png" />
          <img v-if="sbxData[2].level >= 3 && sbxData[2].level < 9" alt="NFT" src="~@/assets/boss/ordinary31.png" />
          <img v-if="sbxData[2].level >= 9" alt="NFT" src="~@/assets/boss/ordinary32.png" />
          <div class="star"><span class="text-stroke">{{ sbxData[2].star }}</span></div>
          <div class="lv"><span class="text-stroke">Lv.{{ sbxData[2].level }}</span></div>
          <div class="redpoint" v-if="redpoint.equipLv3 || redpoint.equipStar3"></div>
        </div>
        <div class="pic baseitem" @click="onSbx(sbxData[3].equip_type)">
          <img v-if="sbxData[3].level < 3" alt="NFT" src="~@/assets/boss/ordinary4.png" />
          <img v-if="sbxData[3].level >= 3 && sbxData[3].level < 9" alt="NFT" src="~@/assets/boss/ordinary41.png" />
          <img v-if="sbxData[3].level >= 9" alt="NFT" src="~@/assets/boss/ordinary42.png" />
          <div class="star"><span class="text-stroke">{{ sbxData[3].star }}</span></div>
          <div class="lv"><span class="text-stroke">Lv.{{ sbxData[3].level }}</span></div>
          <div class="redpoint" v-if="redpoint.equipLv4 || redpoint.equipStar4"></div>
        </div>
        <div class="pic baseitem" @click="onSbx(sbxData[4].equip_type)">
          <img v-if="sbxData[4].level < 3" alt="NFT" src="~@/assets/boss/ordinary5.png" />
          <img v-if="sbxData[4].level >= 3 && sbxData[4].level < 9" alt="NFT" src="~@/assets/boss/ordinary51.png" />
          <img v-if="sbxData[4].level >= 9" alt="NFT" src="~@/assets/boss/ordinary52.png" />
          <div class="star"><span class="text-stroke">{{ sbxData[4].star }}</span></div>
          <div class="lv"><span class="text-stroke">Lv.{{ sbxData[4].level }}</span></div>
          <div class="redpoint" v-if="redpoint.equipLv5 || redpoint.equipStar5"></div>
        </div>
        <div class="user-battle-info" @click="onPopBox(8)">
          <img src="../../assets/info.png" />
        </div>
      </div>
      <div class="bbx">
        <div>
          <div class="pic" @click="openNftEquip(1)">
            <img alt="add"  v-if="nftEquip1 == 0" class="add-img" src="~@/assets/add.png"/>
            <img alt="NFT" v-if="nftEquip1 > 0 && this.nftsEquiped1.level == 1 && this.nftsEquiped1.type == 1" src="~@/assets/nfts/Garm1.png" />
            <img alt="NFT" v-if="nftEquip1 > 0 && this.nftsEquiped1.level == 1 && this.nftsEquiped1.type == 2" src="~@/assets/nfts/Edda1.png" />
            <img alt="NFT" v-if="nftEquip1 > 0 && this.nftsEquiped1.level == 1 && this.nftsEquiped1.type == 3" src="~@/assets/nfts/Norns1.png" />
            <img alt="NFT" v-if="nftEquip1 > 0 && this.nftsEquiped1.level == 1 && this.nftsEquiped1.type == 4" src="~@/assets/nfts/Parvati1.png" />
            <img alt="NFT" v-if="nftEquip1 > 0 && this.nftsEquiped1.level == 1 && this.nftsEquiped1.type == 5" src="~@/assets/nfts/Tartarus1.png" />

            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 1 && this.nftsEquiped1.level <= 3 && this.nftsEquiped1.type == 1" alt="NFT" src="~@/assets/nfts/Garm2.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 1 && this.nftsEquiped1.level <= 3 && this.nftsEquiped1.type == 2" alt="NFT" src="~@/assets/nfts/Edda2.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 1 && this.nftsEquiped1.level <= 3 && this.nftsEquiped1.type == 3" alt="NFT" src="~@/assets/nfts/Norns2.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 1 && this.nftsEquiped1.level <= 3 && this.nftsEquiped1.type == 4" alt="NFT" src="~@/assets/nfts/Parvati2.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 1 && this.nftsEquiped1.level <= 3 && this.nftsEquiped1.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus2.png" />

            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 3 && this.nftsEquiped1.level <= 6 && this.nftsEquiped1.type == 1" alt="NFT" src="~@/assets/nfts/Garm3.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 3 && this.nftsEquiped1.level <= 6 && this.nftsEquiped1.type == 2" alt="NFT" src="~@/assets/nfts/Edda3.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 3 && this.nftsEquiped1.level <= 6 && this.nftsEquiped1.type == 3" alt="NFT" src="~@/assets/nfts/Norns3.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 3 && this.nftsEquiped1.level <= 6 && this.nftsEquiped1.type == 4" alt="NFT" src="~@/assets/nfts/Parvati3.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 3 && this.nftsEquiped1.level <= 6 && this.nftsEquiped1.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus3.png" />

            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 6 && this.nftsEquiped1.level <= 9 && this.nftsEquiped1.type == 1" alt="NFT" src="~@/assets/nfts/Garm4.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 6 && this.nftsEquiped1.level <= 9 && this.nftsEquiped1.type == 2" alt="NFT" src="~@/assets/nfts/Edda4.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 6 && this.nftsEquiped1.level <= 9 && this.nftsEquiped1.type == 3" alt="NFT" src="~@/assets/nfts/Norns4.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 6 && this.nftsEquiped1.level <= 9 && this.nftsEquiped1.type == 4" alt="NFT" src="~@/assets/nfts/Parvati4.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 6 && this.nftsEquiped1.level <= 9 && this.nftsEquiped1.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus4.png" />

            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 9 && this.nftsEquiped1.level <= 12 && this.nftsEquiped1.type == 1" alt="NFT" src="~@/assets/nfts/Garm5.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 9 && this.nftsEquiped1.level <= 12 && this.nftsEquiped1.type == 2" alt="NFT" src="~@/assets/nfts/Edda5.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 9 && this.nftsEquiped1.level <= 12 && this.nftsEquiped1.type == 3" alt="NFT" src="~@/assets/nfts/Norns5.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 9 && this.nftsEquiped1.level <= 12 && this.nftsEquiped1.type == 4" alt="NFT" src="~@/assets/nfts/Parvati5.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 9 && this.nftsEquiped1.level <= 12 && this.nftsEquiped1.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus5.png" />

            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 12 && this.nftsEquiped1.type == 1" alt="NFT" src="~@/assets/nfts/Garm6.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 12 && this.nftsEquiped1.type == 2" alt="NFT" src="~@/assets/nfts/Edda6.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 12 && this.nftsEquiped1.type == 3" alt="NFT" src="~@/assets/nfts/Norns6.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 12 && this.nftsEquiped1.type == 4" alt="NFT" src="~@/assets/nfts/Parvati6.png" />
            <img v-if="nftEquip1 > 0 && this.nftsEquiped1.level > 12 && this.nftsEquiped1.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus6.png" />

            <div class="hgg" v-if="nftEquip1 > 0"><span class="text-stroke">+{{ nftsEquiped1.level }}</span></div>
            <div class="redpoint" v-if="nftEquip1 == 0"></div>
          </div>

          <div class="pic" @click="openNftEquip(2)">

            <img alt="add"  v-if="nftEquip2 == 0" class="add-img" src="~@/assets/add.png"/>
            <img alt="NFT" v-if="nftEquip2 > 0 && this.nftsEquiped2.level == 1 && this.nftsEquiped2.type == 1" src="~@/assets/nfts/Garm1.png" />
            <img alt="NFT" v-if="nftEquip2 > 0 && this.nftsEquiped2.level == 1 && this.nftsEquiped2.type == 2" src="~@/assets/nfts/Edda1.png" />
            <img alt="NFT" v-if="nftEquip2 > 0 && this.nftsEquiped2.level == 1 && this.nftsEquiped2.type == 3" src="~@/assets/nfts/Norns1.png" />
            <img alt="NFT" v-if="nftEquip2 > 0 && this.nftsEquiped2.level == 1 && this.nftsEquiped2.type == 4" src="~@/assets/nfts/Parvati1.png" />
            <img alt="NFT" v-if="nftEquip2 > 0 && this.nftsEquiped2.level == 1 && this.nftsEquiped2.type == 5" src="~@/assets/nfts/Tartarus1.png" />

            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 1 && this.nftsEquiped2.level <= 3 && this.nftsEquiped2.type == 1" alt="NFT" src="~@/assets/nfts/Garm2.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 1 && this.nftsEquiped2.level <= 3 && this.nftsEquiped2.type == 2" alt="NFT" src="~@/assets/nfts/Edda2.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 1 && this.nftsEquiped2.level <= 3 && this.nftsEquiped2.type == 3" alt="NFT" src="~@/assets/nfts/Norns2.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 1 && this.nftsEquiped2.level <= 3 && this.nftsEquiped2.type == 4" alt="NFT" src="~@/assets/nfts/Parvati2.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 1 && this.nftsEquiped2.level <= 3 && this.nftsEquiped2.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus2.png" />

            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 3 && this.nftsEquiped2.level <= 6 && this.nftsEquiped2.type == 1" alt="NFT" src="~@/assets/nfts/Garm3.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 3 && this.nftsEquiped2.level <= 6 && this.nftsEquiped2.type == 2" alt="NFT" src="~@/assets/nfts/Edda3.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 3 && this.nftsEquiped2.level <= 6 && this.nftsEquiped2.type == 3" alt="NFT" src="~@/assets/nfts/Norns3.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 3 && this.nftsEquiped2.level <= 6 && this.nftsEquiped2.type == 4" alt="NFT" src="~@/assets/nfts/Parvati3.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 3 && this.nftsEquiped2.level <= 6 && this.nftsEquiped2.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus3.png" />

            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 6 && this.nftsEquiped2.level <= 9 && this.nftsEquiped2.type == 1" alt="NFT" src="~@/assets/nfts/Garm4.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 6 && this.nftsEquiped2.level <= 9 && this.nftsEquiped2.type == 2" alt="NFT" src="~@/assets/nfts/Edda4.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 6 && this.nftsEquiped2.level <= 9 && this.nftsEquiped2.type == 3" alt="NFT" src="~@/assets/nfts/Norns4.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 6 && this.nftsEquiped2.level <= 9 && this.nftsEquiped2.type == 4" alt="NFT" src="~@/assets/nfts/Parvati4.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 6 && this.nftsEquiped2.level <= 9 && this.nftsEquiped2.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus4.png" />

            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 9 && this.nftsEquiped2.level <= 12 && this.nftsEquiped2.type == 1" alt="NFT" src="~@/assets/nfts/Garm5.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 9 && this.nftsEquiped2.level <= 12 && this.nftsEquiped2.type == 2" alt="NFT" src="~@/assets/nfts/Edda5.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 9 && this.nftsEquiped2.level <= 12 && this.nftsEquiped2.type == 3" alt="NFT" src="~@/assets/nfts/Norns5.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 9 && this.nftsEquiped2.level <= 12 && this.nftsEquiped2.type == 4" alt="NFT" src="~@/assets/nfts/Parvati5.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 9 && this.nftsEquiped2.level <= 12 && this.nftsEquiped2.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus5.png" />

            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 12 && this.nftsEquiped2.type == 1" alt="NFT" src="~@/assets/nfts/Garm6.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 12 && this.nftsEquiped2.type == 2" alt="NFT" src="~@/assets/nfts/Edda6.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 12 && this.nftsEquiped2.type == 3" alt="NFT" src="~@/assets/nfts/Norns6.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 12 && this.nftsEquiped2.type == 4" alt="NFT" src="~@/assets/nfts/Parvati6.png" />
            <img v-if="nftEquip2 > 0 && this.nftsEquiped2.level > 12 && this.nftsEquiped2.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus6.png" />

            <div class="hgg" v-if="nftEquip2 > 0"><span class="text-stroke">+{{ nftsEquiped2.level }}</span></div>
            <div class="redpoint" v-if="nftEquip2 == 0"></div>
          </div>
          <div class="pic" @click="openNftEquip(3)">
            <img alt="add"  v-if="nftEquip3 == 0" class="add-img" src="~@/assets/add.png"/>
            <img alt="NFT" v-if="nftEquip3 > 0 && this.nftsEquiped3.level == 1 && this.nftsEquiped3.type == 1" src="~@/assets/nfts/Garm1.png" />
            <img alt="NFT" v-if="nftEquip3 > 0 && this.nftsEquiped3.level == 1 && this.nftsEquiped3.type == 2" src="~@/assets/nfts/Edda1.png" />
            <img alt="NFT" v-if="nftEquip3 > 0 && this.nftsEquiped3.level == 1 && this.nftsEquiped3.type == 3" src="~@/assets/nfts/Norns1.png" />
            <img alt="NFT" v-if="nftEquip3 > 0 && this.nftsEquiped3.level == 1 && this.nftsEquiped3.type == 4" src="~@/assets/nfts/Parvati1.png" />
            <img alt="NFT" v-if="nftEquip3 > 0 && this.nftsEquiped3.level == 1 && this.nftsEquiped3.type == 5" src="~@/assets/nfts/Tartarus1.png" />

            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 1 && this.nftsEquiped3.level <= 3 && this.nftsEquiped3.type == 1" alt="NFT" src="~@/assets/nfts/Garm2.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 1 && this.nftsEquiped3.level <= 3 && this.nftsEquiped3.type == 2" alt="NFT" src="~@/assets/nfts/Edda2.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 1 && this.nftsEquiped3.level <= 3 && this.nftsEquiped3.type == 3" alt="NFT" src="~@/assets/nfts/Norns2.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 1 && this.nftsEquiped3.level <= 3 && this.nftsEquiped3.type == 4" alt="NFT" src="~@/assets/nfts/Parvati2.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 1 && this.nftsEquiped3.level <= 3 && this.nftsEquiped3.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus2.png" />

            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 3 && this.nftsEquiped3.level <= 6 && this.nftsEquiped3.type == 1" alt="NFT" src="~@/assets/nfts/Garm3.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 3 && this.nftsEquiped3.level <= 6 && this.nftsEquiped3.type == 2" alt="NFT" src="~@/assets/nfts/Edda3.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 3 && this.nftsEquiped3.level <= 6 && this.nftsEquiped3.type == 3" alt="NFT" src="~@/assets/nfts/Norns3.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 3 && this.nftsEquiped3.level <= 6 && this.nftsEquiped3.type == 4" alt="NFT" src="~@/assets/nfts/Parvati3.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 3 && this.nftsEquiped3.level <= 6 && this.nftsEquiped3.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus3.png" />

            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 6 && this.nftsEquiped3.level <= 9 && this.nftsEquiped3.type == 1" alt="NFT" src="~@/assets/nfts/Garm4.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 6 && this.nftsEquiped3.level <= 9 && this.nftsEquiped3.type == 2" alt="NFT" src="~@/assets/nfts/Edda4.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 6 && this.nftsEquiped3.level <= 9 && this.nftsEquiped3.type == 3" alt="NFT" src="~@/assets/nfts/Norns4.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 6 && this.nftsEquiped3.level <= 9 && this.nftsEquiped3.type == 4" alt="NFT" src="~@/assets/nfts/Parvati4.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 6 && this.nftsEquiped3.level <= 9 && this.nftsEquiped3.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus4.png" />

            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 9 && this.nftsEquiped3.level <= 12 && this.nftsEquiped3.type == 1" alt="NFT" src="~@/assets/nfts/Garm5.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 9 && this.nftsEquiped3.level <= 12 && this.nftsEquiped3.type == 2" alt="NFT" src="~@/assets/nfts/Edda5.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 9 && this.nftsEquiped3.level <= 12 && this.nftsEquiped3.type == 3" alt="NFT" src="~@/assets/nfts/Norns5.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 9 && this.nftsEquiped3.level <= 12 && this.nftsEquiped3.type == 4" alt="NFT" src="~@/assets/nfts/Parvati5.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 9 && this.nftsEquiped3.level <= 12 && this.nftsEquiped3.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus5.png" />

            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 12 && this.nftsEquiped3.type == 1" alt="NFT" src="~@/assets/nfts/Garm6.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 12 && this.nftsEquiped3.type == 2" alt="NFT" src="~@/assets/nfts/Edda6.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 12 && this.nftsEquiped3.type == 3" alt="NFT" src="~@/assets/nfts/Norns6.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 12 && this.nftsEquiped3.type == 4" alt="NFT" src="~@/assets/nfts/Parvati6.png" />
            <img v-if="nftEquip3 > 0 && this.nftsEquiped3.level > 12 && this.nftsEquiped3.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus6.png" />

            <div class="hgg" v-if="nftEquip3 > 0"><span class="text-stroke">+{{ nftsEquiped3.level }}</span></div>
            <div class="redpoint" v-if="nftEquip3 == 0"></div>
          </div>
        </div>

        <div class="hgg-total">
          <span class="text-stroke">{{ nftLevelTotal }}</span>
        </div>
      </div>
    </div>

    <foots />


    <!-- pop start -->
    <div class="popbox mypopbox" v-if="istip">
      <div class="bgs"></div>
      <!-- type 1-->
      <div class="popbody nftbody" v-if="isType === 1">
        <div class="head">
          <span class="text-stroke">Weapon</span>
          <span class="icon"></span>
        </div>
        <div class="middle">
          <div class="type">
            <div v-if="!hasChoicedNFT" class="blockhe">
              <span class="text-stroke">Choose a NFT weapon below</span>
            </div>
            <div v-else class="blockhe">
              <div class="left">
                <div v-if="this.nftChoiced.type == 1" class="picpkc">
                  <img v-if="this.nftChoiced.level == 1" alt="NFT" src="~@/assets/nfts/Garm1.png" />
                  <img v-if="this.nftChoiced.level > 1 && this.nftChoiced.level <= 3" alt="NFT" src="~@/assets/nfts/Garm2.png" />
                  <img v-if="this.nftChoiced.level > 3 && this.nftChoiced.level <= 6" alt="NFT" src="~@/assets/nfts/Garm3.png" />
                  <img v-if="this.nftChoiced.level > 6 && this.nftChoiced.level <= 9" alt="NFT" src="~@/assets/nfts/Garm4.png" />
                  <img v-if="this.nftChoiced.level > 9 && this.nftChoiced.level <= 12" alt="NFT" src="~@/assets/nfts/Garm5.png" />
                  <img v-if="this.nftChoiced.level > 12" alt="NFT" src="~@/assets/nfts/Garm6.png" />
                  <div class="hggiconss"><span class="text-stroke">+{{ this.nftChoiced.level }}</span></div>
                </div>
                <div v-if="this.nftChoiced.type == 2" class="picpkc">
                  <img v-if="this.nftChoiced.level == 1" alt="NFT" src="~@/assets/nfts/Edda1.png" />
                  <img v-if="this.nftChoiced.level > 1 && this.nftChoiced.level <= 3" alt="NFT" src="~@/assets/nfts/Edda2.png" />
                  <img v-if="this.nftChoiced.level > 3 && this.nftChoiced.level <= 6" alt="NFT" src="~@/assets/nfts/Edda3.png" />
                  <img v-if="this.nftChoiced.level > 6 && this.nftChoiced.level <= 9" alt="NFT" src="~@/assets/nfts/Edda4.png" />
                  <img v-if="this.nftChoiced.level > 9 && this.nftChoiced.level <= 12" alt="NFT" src="~@/assets/nfts/Edda5.png" />
                  <img v-if="this.nftChoiced.level > 12" alt="NFT" src="~@/assets/nfts/Edda6.png" />
                  <div class="hggiconss"><span class="text-stroke">+{{ this.nftChoiced.level }}</span></div>
                </div>
                <div v-if="this.nftChoiced.type == 3" class="picpkc">
                  <img v-if="this.nftChoiced.level == 1" alt="NFT" src="~@/assets/nfts/Norns1.png" />
                  <img v-if="this.nftChoiced.level > 1 && this.nftChoiced.level <= 3" alt="NFT" src="~@/assets/nfts/Norns2.png" />
                  <img v-if="this.nftChoiced.level > 3 && this.nftChoiced.level <= 6" alt="NFT" src="~@/assets/nfts/Norns3.png" />
                  <img v-if="this.nftChoiced.level > 6 && this.nftChoiced.level <= 9" alt="NFT" src="~@/assets/nfts/Norns4.png" />
                  <img v-if="this.nftChoiced.level > 9 && this.nftChoiced.level <= 12" alt="NFT" src="~@/assets/nfts/Norns5.png" />
                  <img v-if="this.nftChoiced.level > 12" alt="NFT" src="~@/assets/nfts/Norns6.png" />
                  <div class="hggiconss"><span class="text-stroke">+{{ this.nftChoiced.level }}</span></div>
                </div>
                <div v-if="this.nftChoiced.type == 4" class="picpkc">
                  <img v-if="this.nftChoiced.level == 1" alt="NFT" src="~@/assets/nfts/Parvati1.png" />
                  <img v-if="this.nftChoiced.level > 1 && this.nftChoiced.level <= 3" alt="NFT" src="~@/assets/nfts/Parvati2.png" />
                  <img v-if="this.nftChoiced.level > 3 && this.nftChoiced.level <= 6" alt="NFT" src="~@/assets/nfts/Parvati3.png" />
                  <img v-if="this.nftChoiced.level > 6 && this.nftChoiced.level <= 9" alt="NFT" src="~@/assets/nfts/Parvati4.png" />
                  <img v-if="this.nftChoiced.level > 9 && this.nftChoiced.level <= 12" alt="NFT" src="~@/assets/nfts/Parvati5.png" />
                  <img v-if="this.nftChoiced.level > 12" alt="NFT" src="~@/assets/nfts/Parvati6.png" />
                  <div class="hggiconss"><span class="text-stroke">+{{ this.nftChoiced.level }}</span></div>
                </div>
                <div v-if="this.nftChoiced.type == 5" class="picpkc">
                  <img v-if="this.nftChoiced.level == 1" alt="NFT" src="~@/assets/nfts/Tartarus1.png" />
                  <img v-if="this.nftChoiced.level > 1 && this.nftChoiced.level <= 3" alt="NFT" src="~@/assets/nfts/Tartarus2.png" />
                  <img v-if="this.nftChoiced.level > 3 && this.nftChoiced.level <= 6" alt="NFT" src="~@/assets/nfts/Tartarus3.png" />
                  <img v-if="this.nftChoiced.level > 6 && this.nftChoiced.level <= 9" alt="NFT" src="~@/assets/nfts/Tartarus4.png" />
                  <img v-if="this.nftChoiced.level > 9 && this.nftChoiced.level <= 12" alt="NFT" src="~@/assets/nfts/Tartarus5.png" />
                  <img v-if="this.nftChoiced.level > 12" alt="NFT" src="~@/assets/nfts/Tartarus6.png" />
                  <div class="hggiconss"><span class="text-stroke">+{{ this.nftChoiced.level }}</span></div>
                </div>
              </div>
              <div class="right">
                <div v-if="this.nftChoiced.id == this.nftEquip1 || this.nftChoiced.id == this.nftEquip2 || this.nftChoiced.id == this.nftEquip3" class="equipstate text-stroke">Equiped</div>
                <div class="tlsd text-stroke">
                  <span class="text-stroke">{{this.nftChoiced.name}}</span>
                  <span class="nft-id text-stroke">#Test</span>
                </div>
                <div class="atr nft-atr">
                  <div>
                    <span class="text-stroke">{{this.nftChoiced.main_atr_1}}: {{this.nftChoiced.main_atr_1_val / 100}}%</span>
                    <span class="text-stroke">{{this.nftChoiced.main_atr_2}}: {{this.nftChoiced.main_atr_2_val / 100}}%</span>
                  </div>
                  <div>
                    <span v-if="this.nftChoiced.add_atr_1_val > 0" class="text-stroke">{{this.nftChoiced.add_atr_1}}: {{this.nftChoiced.add_atr_1_val / 100}}%</span>
                    <span v-if="this.nftChoiced.add_atr_2_val > 0" class="text-stroke">{{this.nftChoiced.add_atr_2}}: {{this.nftChoiced.add_atr_2_val / 100}}%</span>
                  </div>
                </div>
              </div>
            </div>
            <div class="nft-list">
              <div v-for="item in this.nftsListed" :key="item.id">
                <div v-if="item.type == 1" class="picpkc" @click="choiceNft(item.id)">
                  <img v-if="item.id == nftChoiced.id" class="choiced" src="../../assets/icon-border.png"/>
                  <img v-if="item.level == 1" alt="NFT" src="~@/assets/nfts/Garm1.png" />
                  <img v-if="item.level > 1 && item.level <= 3" alt="NFT" src="~@/assets/nfts/Garm2.png" />
                  <img v-if="item.level > 3 && item.level <= 6" alt="NFT" src="~@/assets/nfts/Garm3.png" />
                  <img v-if="item.level > 6 && item.level <= 9" alt="NFT" src="~@/assets/nfts/Garm4.png" />
                  <img v-if="item.level > 9 && item.level <= 12" alt="NFT" src="~@/assets/nfts/Garm5.png" />
                  <img v-if="item.level > 12" alt="NFT" src="~@/assets/nfts/Garm6.png" />
                  <div class="hggiconss"><span class="text-stroke">+{{ item.level }}</span></div>
                  <div v-if="item.id == nftEquip1 || item.id == nftEquip2 || item.id == nftEquip3" class="equipicon"><span class="text-stroke">E</span></div>
                </div>
                <div v-if="item.type == 2" class="picpkc" @click="choiceNft(item.id)">
                  <img v-if="item.id == nftChoiced.id" class="choiced" src="../../assets/icon-border.png"/>
                  <img v-if="item.level == 1" alt="NFT" src="~@/assets/nfts/Edda1.png" />
                  <img v-if="item.level > 1 && item.level <= 3" alt="NFT" src="~@/assets/nfts/Edda2.png" />
                  <img v-if="item.level > 3 && item.level <= 6" alt="NFT" src="~@/assets/nfts/Edda3.png" />
                  <img v-if="item.level > 6 && item.level <= 9" alt="NFT" src="~@/assets/nfts/Edda4.png" />
                  <img v-if="item.level > 9 && item.level <= 12" alt="NFT" src="~@/assets/nfts/Edda5.png" />
                  <img v-if="item.level > 12" alt="NFT" src="~@/assets/nfts/Edda6.png" />
                  <div class="hggiconss"><span class="text-stroke">+{{ item.level }}</span></div>
                  <div v-if="item.id == nftEquip1 || item.id == nftEquip2 || item.id == nftEquip3" class="equipicon"><span class="text-stroke">E</span></div>
                </div>
                <div v-if="item.type == 3" class="picpkc" @click="choiceNft(item.id)">
                  <img v-if="item.id == nftChoiced.id" class="choiced" src="../../assets/icon-border.png"/>
                  <img v-if="item.level == 1" alt="NFT" src="~@/assets/nfts/Norns1.png" />
                  <img v-if="item.level > 1 && item.level <= 3" alt="NFT" src="~@/assets/nfts/Norns2.png" />
                  <img v-if="item.level > 3 && item.level <= 6" alt="NFT" src="~@/assets/nfts/Norns3.png" />
                  <img v-if="item.level > 6 && item.level <= 9" alt="NFT" src="~@/assets/nfts/Norns4.png" />
                  <img v-if="item.level > 9 && item.level <= 12" alt="NFT" src="~@/assets/nfts/Norns5.png" />
                  <img v-if="item.level > 12" alt="NFT" src="~@/assets/nfts/Norns6.png" />
                  <div class="hggiconss"><span class="text-stroke">+{{ item.level }}</span></div>
                  <div v-if="item.id == nftEquip1 || item.id == nftEquip2 || item.id == nftEquip3" class="equipicon"><span class="text-stroke">E</span></div>
                </div>
                <div v-if="item.type == 4" class="picpkc" @click="choiceNft(item.id)">
                  <img v-if="item.id == nftChoiced.id" class="choiced" src="../../assets/icon-border.png"/>
                  <img v-if="item.level == 1" alt="NFT" src="~@/assets/nfts/Parvati1.png" />
                  <img v-if="item.level > 1 && item.level <= 3" alt="NFT" src="~@/assets/nfts/Parvati2.png" />
                  <img v-if="item.level > 3 && item.level <= 6" alt="NFT" src="~@/assets/nfts/Parvati3.png" />
                  <img v-if="item.level > 6 && item.level <= 9" alt="NFT" src="~@/assets/nfts/Parvati4.png" />
                  <img v-if="item.level > 9 && item.level <= 12" alt="NFT" src="~@/assets/nfts/Parvati5.png" />
                  <img v-if="item.level > 12" alt="NFT" src="~@/assets/nfts/Parvati6.png" />
                  <div class="hggiconss"><span class="text-stroke">+{{ item.level }}</span></div>
                  <div v-if="item.id == nftEquip1 || item.id == nftEquip2 || item.id == nftEquip3" class="equipicon"><span class="text-stroke">E</span></div>
                </div>
                <div v-if="item.type == 5" class="picpkc" @click="choiceNft(item.id)">
                  <img v-if="item.id == nftChoiced.id" class="choiced" src="../../assets/icon-border.png"/>
                  <img v-if="item.level == 1" alt="NFT" src="~@/assets/nfts/Tartarus1.png" />
                  <img v-if="item.level > 1 && item.level <= 3" alt="NFT" src="~@/assets/nfts/Tartarus2.png" />
                  <img v-if="item.level > 3 && item.level <= 6" alt="NFT" src="~@/assets/nfts/Tartarus3.png" />
                  <img v-if="item.level > 6 && item.level <= 9" alt="NFT" src="~@/assets/nfts/Tartarus4.png" />
                  <img v-if="item.level > 9 && item.level <= 12" alt="NFT" src="~@/assets/nfts/Tartarus5.png" />
                  <img v-if="item.level > 12" alt="NFT" src="~@/assets/nfts/Tartarus6.png" />
                  <div class="hggiconss"><span class="text-stroke">+{{ item.level }}</span></div>
                  <div v-if="item.id == nftEquip1 || item.id == nftEquip2 || item.id == nftEquip3" class="equipicon"><span class="text-stroke">E</span></div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="foots" style="display: flex;">
          <div v-if="this.nftChoiced.id == this.nftEquip1 || this.nftChoiced.id == this.nftEquip2 || this.nftChoiced.id == this.nftEquip3" class="btnss-unable f30" style="margin-top: 20px;"><span class="text-stroke">Equiped</span></div>
          <div v-else class="btnss f30" style="margin-top: 20px;" @click="equipNft()"><span class="text-stroke">Equip</span></div>
          <div class="btnss f30" style="margin-top: 20px;" @click="onOpenNftUp(1)"><span class="text-stroke">Upgrade</span></div>
        </div>
        <div class="close-btn" @click="onPopClose()"></div>
      </div>


      <!-- type 2-->
      <!-- NFT upgrade -->
      <div class="popbody nftpop" v-if="isType === 2">
        <div class="head">
          <span class="text-stroke">{{this.nftChoiced.name}}</span>
          <span class="icon"></span>
        </div>
        <div class="middle">
          <div class="type">
            <div class="picpkc" v-if="this.nftChoiced.type === 1">
              <img v-if="this.nftChoiced.level == 1" alt="NFT" src="~@/assets/nfts/Garm1.png" />
              <img v-if="this.nftChoiced.level > 1 && this.nftChoiced.level <= 3" alt="NFT" src="~@/assets/nfts/Garm2.png" />
              <img v-if="this.nftChoiced.level > 3 && this.nftChoiced.level <= 6" alt="NFT" src="~@/assets/nfts/Garm3.png" />
              <img v-if="this.nftChoiced.level > 6 && this.nftChoiced.level <= 9" alt="NFT" src="~@/assets/nfts/Garm4.png" />
              <img v-if="this.nftChoiced.level > 9 && this.nftChoiced.level <= 12" alt="NFT" src="~@/assets/nfts/Garm5.png" />
              <img v-if="this.nftChoiced.level > 12" alt="NFT" src="~@/assets/nfts/Garm6.png" />
            </div>
            <div class="picpkc" v-if="this.nftChoiced.type === 2">
              <img v-if="this.nftChoiced.level == 1" alt="NFT" src="~@/assets/nfts/Edda1.png" />
              <img v-if="this.nftChoiced.level > 1 && this.nftChoiced.level <= 3" alt="NFT" src="~@/assets/nfts/Edda2.png" />
              <img v-if="this.nftChoiced.level > 3 && this.nftChoiced.level <= 6" alt="NFT" src="~@/assets/nfts/Edda3.png" />
              <img v-if="this.nftChoiced.level > 6 && this.nftChoiced.level <= 9" alt="NFT" src="~@/assets/nfts/Edda4.png" />
              <img v-if="this.nftChoiced.level > 9 && this.nftChoiced.level <= 12" alt="NFT" src="~@/assets/nfts/Edda5.png" />
              <img v-if="this.nftChoiced.level > 12" alt="NFT" src="~@/assets/nfts/Edda6.png" />
            </div>
            <div class="picpkc" v-if="this.nftChoiced.type === 3">
              <img v-if="this.nftChoiced.level == 1" alt="NFT" src="~@/assets/nfts/Norns1.png" />
              <img v-if="this.nftChoiced.level > 1 && this.nftChoiced.level <= 3" alt="NFT" src="~@/assets/nfts/Norns2.png" />
              <img v-if="this.nftChoiced.level > 3 && this.nftChoiced.level <= 6" alt="NFT" src="~@/assets/nfts/Norns3.png" />
              <img v-if="this.nftChoiced.level > 6 && this.nftChoiced.level <= 9" alt="NFT" src="~@/assets/nfts/Norns4.png" />
              <img v-if="this.nftChoiced.level > 9 && this.nftChoiced.level <= 12" alt="NFT" src="~@/assets/nfts/Norns5.png" />
              <img v-if="this.nftChoiced.level > 12" alt="NFT" src="~@/assets/nfts/Norns6.png" />
            </div>
            <div class="picpkc" v-if="this.nftChoiced.type === 4">
              <img v-if="this.nftChoiced.level == 1" alt="NFT" src="~@/assets/nfts/Parvati1.png" />
              <img v-if="this.nftChoiced.level > 1 && this.nftChoiced.level <= 3" alt="NFT" src="~@/assets/nfts/Parvati2.png" />
              <img v-if="this.nftChoiced.level > 3 && this.nftChoiced.level <= 6" alt="NFT" src="~@/assets/nfts/Parvati3.png" />
              <img v-if="this.nftChoiced.level > 6 && this.nftChoiced.level <= 9" alt="NFT" src="~@/assets/nfts/Parvati4.png" />
              <img v-if="this.nftChoiced.level > 9 && this.nftChoiced.level <= 12" alt="NFT" src="~@/assets/nfts/Parvati5.png" />
              <img v-if="this.nftChoiced.level > 12" alt="NFT" src="~@/assets/nfts/Parvati6.png" />
            </div>
            <div class="picpkc" v-if="this.nftChoiced.type === 5">
              <img v-if="this.nftChoiced.level == 1" alt="NFT" src="~@/assets/nfts/Tartarus1.png" />
              <img v-if="this.nftChoiced.level > 1 && this.nftChoiced.level <= 3" alt="NFT" src="~@/assets/nfts/Tartarus2.png" />
              <img v-if="this.nftChoiced.level > 3 && this.nftChoiced.level <= 6" alt="NFT" src="~@/assets/nfts/Tartarus3.png" />
              <img v-if="this.nftChoiced.level > 6 && this.nftChoiced.level <= 9" alt="NFT" src="~@/assets/nfts/Tartarus4.png" />
              <img v-if="this.nftChoiced.level > 9 && this.nftChoiced.level <= 12" alt="NFT" src="~@/assets/nfts/Tartarus5.png" />
              <img v-if="this.nftChoiced.level > 12" alt="NFT" src="~@/assets/nfts/Tartarus6.png" />
            </div>
            <!-- end 3 -->
            <div class="uphgg text-stroke">
              <em></em>
              <span>{{ this.nftChoiced.level }}</span>
              <span class="upimg" v-if="this.nftChoiced.level < 15"></span>
              <em v-if="this.nftChoiced.level < 15"></em>
              <span v-if="this.nftChoiced.level < 15">{{ this.nftChoiced.level + 1 }}</span>
            </div>
            <div class="atr2 nftatr2 text-stroke" style=" margin:20px auto 0; overflow: hidden; padding-bottom: 30px;">
              <div class="text-stroke" style="width: 100%; text-align: center; text-indent:0; margin-top: 10px;">
                <span :class="showuptxtani ? 'atr-ani' : ''">
                  {{this.nftChoiced.main_atr_1}}: {{ this.nftChoiced.main_atr_1_val / 100 }}%
                </span>
                <span v-if="this.nftChoiced.level < 15" class="atr-add">+1%</span>
              </div>
              <div class="text-stroke" style="width: 100%; text-align: center; text-indent:0; margin-top: 10px;">
                <span :class="showuptxtani ? 'atr-ani' : ''">
                  {{this.nftChoiced.main_atr_2}}: {{ this.nftChoiced.main_atr_2_val / 100 }}%
                </span>
                <span v-if="this.nftChoiced.level < 15" class="atr-add">+1%</span>
              </div>
              <div v-if="this.nftChoiced.add_atr_1_val > 0" class="text-stroke" style="width: 100%; text-align: center; text-indent:0; margin-top: 10px;">
                <span :class="showuptxtani2 ? 'atr-ani' : ''">
                  {{this.nftChoiced.add_atr_1}}: {{ this.nftChoiced.add_atr_1_val / 100 }}%
                </span>
              </div>
              <div v-if="this.nftChoiced.add_atr_2_val > 0" class="text-stroke" style="width: 100%; text-align: center; text-indent:0; margin-top: 10px;">
                <span :class="showuptxtani3 ? 'atr-ani' : ''">
                  {{this.nftChoiced.add_atr_2}}: {{ this.nftChoiced.add_atr_2_val / 100 }}%
                </span>
              </div>
            </div>
            <div class="addAtrInfo">
              <span class="text-stroke">When the Lv. reaches 3/6/9/12/15, a random attribute will be added.</span>
            </div>
          </div>
        </div>
        <div class="foots ">
          <div class="btxthgg">
            <p class="text-stroke">GAS<span> ≈ 0.01 TON</span></p>
            <p class="text-stroke">COST:<img src="../../assets/ETHERicon.png" /><span>1 ≈ 1U</span></p>
          </div>
          <div class="btnss nftupbtn f30" @click="onUpgradeNFT()"><span class="text-stroke">Upgrade</span></div>
        </div>
        <div class="close-btn" @click="onPopClose()"></div>
      </div>

      <!-- type 3 普通武器升级弹框-->
      <div class="popbody" v-if="isType === 3">
        <div class="head">
          <span class="text-stroke">Equip</span>
        </div>
        <div class="middle">
          <div class="type">
            <div class="picpkc baseitem">
              <img alt="NFT" v-if="sbxIndex === 1 && ordinary.level < 3" src="~@/assets/boss/ordinary1.png" />
              <img alt="NFT" v-if="sbxIndex === 1 && ordinary.level >= 3 && ordinary.level < 9" src="~@/assets/boss/ordinary11.png" />
              <img alt="NFT" v-if="sbxIndex === 1 && ordinary.level >= 9" src="~@/assets/boss/ordinary12.png" />
              <img alt="NFT" v-if="sbxIndex === 2 && ordinary.level < 3" src="~@/assets/boss/ordinary2.png" />
              <img alt="NFT" v-if="sbxIndex === 2 && ordinary.level >= 3 && ordinary.level < 9" src="~@/assets/boss/ordinary21.png" />
              <img alt="NFT" v-if="sbxIndex === 2 && ordinary.level >= 9" src="~@/assets/boss/ordinary22.png" />
              <img alt="NFT" v-if="sbxIndex === 3 && ordinary.level < 3" src="~@/assets/boss/ordinary3.png" />
              <img alt="NFT" v-if="sbxIndex === 3 && ordinary.level >= 3 && ordinary.level < 9" src="~@/assets/boss/ordinary31.png" />
              <img alt="NFT" v-if="sbxIndex === 3 && ordinary.level >= 9" src="~@/assets/boss/ordinary32.png" />
              <img alt="NFT" v-if="sbxIndex === 4 && ordinary.level < 3" src="~@/assets/boss/ordinary4.png" />
              <img alt="NFT" v-if="sbxIndex === 4 && ordinary.level >= 3 && ordinary.level < 9" src="~@/assets/boss/ordinary41.png" />
              <img alt="NFT" v-if="sbxIndex === 4 && ordinary.level >=9" src="~@/assets/boss/ordinary42.png" />
              <img alt="NFT" v-if="sbxIndex === 5 && ordinary.level < 3" src="~@/assets/boss/ordinary5.png" />
              <img alt="NFT" v-if="sbxIndex === 5 && ordinary.level >= 3 && ordinary.level < 9" src="~@/assets/boss/ordinary51.png" />
              <img alt="NFT" v-if="sbxIndex === 5 && ordinary.level >= 9" src="~@/assets/boss/ordinary52.png" />

              <div class="star"><span class="text-stroke">{{ ordinary.star }}</span></div>
              <div class="lv"><span class="text-stroke">Lv.{{ ordinary.level }}</span></div>
            </div>
            <div class="atr" style="margin:20px auto 0; overflow: hidden; padding-bottom: 30px;">
              <span class="text-stroke" style="width: 100%; text-align: center; text-indent:0; margin-top: 10px;">HP：{{ ordinary.attrs ?
                ordinary.attrs.hp : 0 }} </span>
              <span class="text-stroke" style="width: 100%; text-align: center; text-indent:0; margin-top: 10px;">ATK：{{ ordinary.attrs ?
                ordinary.attrs.atk : 0 }}
              </span>
            </div>
          </div>
        </div>
        <div class="foots" style="display: flex;">
          <div class="btnss f30" @click="upgradeData(1)">
            <div class="btn-redpoint" v-if="ordinary.equip_type == 1 && redpoint.equipLv1"></div>
            <div class="btn-redpoint" v-if="ordinary.equip_type == 2 && redpoint.equipLv2"></div>
            <div class="btn-redpoint" v-if="ordinary.equip_type == 3 && redpoint.equipLv3"></div>
            <div class="btn-redpoint" v-if="ordinary.equip_type == 4 && redpoint.equipLv4"></div>
            <div class="btn-redpoint" v-if="ordinary.equip_type == 5 && redpoint.equipLv5"></div>
            <span class="text-stroke">Upgrade</span>
          </div>
          <div class="btnss f30" @click="upgradeData(2)">
            <div class="btn-redpoint" v-if="ordinary.equip_type == 1 && redpoint.equipStar1"></div>
            <div class="btn-redpoint" v-if="ordinary.equip_type == 2 && redpoint.equipStar2"></div>
            <div class="btn-redpoint" v-if="ordinary.equip_type == 3 && redpoint.equipStar3"></div>
            <div class="btn-redpoint" v-if="ordinary.equip_type == 4 && redpoint.equipStar4"></div>
            <div class="btn-redpoint" v-if="ordinary.equip_type == 5 && redpoint.equipStar5"></div>
            <span class="text-stroke">Enhance</span>
          </div>
        </div>
        <div class="close-btn" @click="onPopClose()"></div>
      </div>

      <!-- type 4 普通武器升级升星 -->
      <div class="popbody" v-if="isType === 4">
        <div class="head">
          <span class="text-stroke">Equip</span>
        </div>
        <div class="middle">
          <div class="type">
            <div class="picpkc baseitem">
              <div class="upani" v-if="showupani"></div>
              <img alt="NFT" v-if="sbxIndex === 1 && ordinary.level < 3" src="~@/assets/boss/ordinary1.png" />
              <img alt="NFT" v-if="sbxIndex === 1 && ordinary.level >= 3 && ordinary.level < 9" src="~@/assets/boss/ordinary11.png" />
              <img alt="NFT" v-if="sbxIndex === 1 && ordinary.level >= 9" src="~@/assets/boss/ordinary12.png" />
              <img alt="NFT" v-if="sbxIndex === 2 && ordinary.level < 3" src="~@/assets/boss/ordinary2.png" />
              <img alt="NFT" v-if="sbxIndex === 2 && ordinary.level >= 3 && ordinary.level < 9" src="~@/assets/boss/ordinary21.png" />
              <img alt="NFT" v-if="sbxIndex === 2 && ordinary.level >= 9" src="~@/assets/boss/ordinary22.png" />
              <img alt="NFT" v-if="sbxIndex === 3 && ordinary.level < 3" src="~@/assets/boss/ordinary3.png" />
              <img alt="NFT" v-if="sbxIndex === 3 && ordinary.level >= 3 && ordinary.level < 9" src="~@/assets/boss/ordinary31.png" />
              <img alt="NFT" v-if="sbxIndex === 3 && ordinary.level >= 9" src="~@/assets/boss/ordinary32.png" />
              <img alt="NFT" v-if="sbxIndex === 4 && ordinary.level < 3" src="~@/assets/boss/ordinary4.png" />
              <img alt="NFT" v-if="sbxIndex === 4 && ordinary.level >= 3 && ordinary.level < 9" src="~@/assets/boss/ordinary41.png" />
              <img alt="NFT" v-if="sbxIndex === 4 && ordinary.level >=9" src="~@/assets/boss/ordinary42.png" />
              <img alt="NFT" v-if="sbxIndex === 5 && ordinary.level < 3" src="~@/assets/boss/ordinary5.png" />
              <img alt="NFT" v-if="sbxIndex === 5 && ordinary.level >= 3 && ordinary.level < 9" src="~@/assets/boss/ordinary51.png" />
              <img alt="NFT" v-if="sbxIndex === 5 && ordinary.level >= 9" src="~@/assets/boss/ordinary52.png" />
            </div>
            <div class="upvl text-stroke" v-if="isUpgrade === 1">
              <span>Lv.{{ ordinary.level }}</span>
              <span class="upimg" v-if="ordinary.level < 20"></span>
              <span v-if="ordinary.level < 20">Lv.{{ ordinary.level + 1 }}</span>
            </div>
            <div class="upstar text-stroke" v-if="isUpgrade === 2">
              <em></em>
              <span>{{ ordinary.star }}</span>
              <span class="upimg" v-if="ordinary.star < 20"></span>
              <em v-if="ordinary.star < 20"></em>
              <span v-if="ordinary.star < 20">{{ ordinary.star + 1 }}</span>
            </div>

            <div class="atr2" v-if="isUpgrade === 1" style="margin:20px auto 0; overflow: hidden; padding-bottom: 30px;">
              <div class="text-stroke" style="width: 100%; text-align: center; text-indent:0; margin-top: 10px;">
                <span>HP:</span>
                <span :class="showuptxtani ? 'atr-ani' : ''">{{ ordinary.attrs ? ordinary.attrs.hp : 0 }}</span>
                <span v-if="ordinary.level < 20" class="atr-add">+{{upAddData.baseLvHpAdd}}</span>
              </div>
              <div class="text-stroke" style="width: 100%; text-align: center; text-indent:0; margin-top: 10px;">
                <span>ATK:</span>
                <span :class="showuptxtani ? 'atr-ani' : ''">{{ ordinary.attrs ? ordinary.attrs.atk : 0 }}</span>
                <span v-if="ordinary.level < 20" class="atr-add">+{{upAddData.baseLvAtkAdd}}</span>
              </div>
            </div>
            <div class="atr2" v-if="isUpgrade === 2" style="margin:20px auto 0; overflow: hidden; padding-bottom: 30px;">
              <div class="text-stroke" style="width: 100%; text-align: center; text-indent:0; margin-top: 10px;">
                <span>HP:</span>
                <span :class="showuptxtani ? 'atr-ani' : ''">{{ ordinary.attrs ? ordinary.attrs.hp : 0 }}</span>
                <span v-if="ordinary.star < 20" class="atr-add">+{{upAddData.baseStarHpAdd}}</span>
              </div>
              <div class="text-stroke" style="width: 100%; text-align: center; text-indent:0; margin-top: 10px;">
                <span>ATK:</span>
                <span :class="showuptxtani ? 'atr-ani' : ''">{{ ordinary.attrs ? ordinary.attrs.atk : 0 }}</span>
                <span v-if="ordinary.star < 20" class="atr-add">+{{upAddData.baseStarAtkAdd}}</span>
              </div>
            </div>
            <div class="uprate text-stroke" v-if="isUpgrade === 1 && ordinary.level < 20">
              <span v-if="baseEquipLevelData[ordinary.equip_type + '_' + ordinary.level].UpgradeSuccessRate < 10000">
                {{upAddData.uplvFailedNum}}/{{baseEquipLevelData[ordinary.equip_type + "_" + ordinary.level].ProtectSuccessNum}} times is sure to succeed
              </span>
              <span v-if="upAddData.uplvFailedNum != baseEquipLevelData[ordinary.equip_type + '_' + ordinary.level].ProtectSuccessNum" :class="baseEquipLevelData[ordinary.equip_type + '_' + ordinary.level].UpgradeSuccessRate < 10000 ? '' : 'success'">
                Success rate: {{baseEquipLevelData[ordinary.equip_type + "_" + ordinary.level].UpgradeSuccessRate/100}}%
              </span>
              <span v-else class="success">
                Success rate: 100%
              </span>
            </div>
            <div class="uprate text-stroke" v-if="isUpgrade === 2 && ordinary.star < 20">
              <span v-if="baseEquipStarData[ordinary.equip_type + '_' + ordinary.star].UpgradeSuccessRate < 10000">
                {{upAddData.upstarFailedNum}}/{{baseEquipStarData[ordinary.equip_type + "_" + ordinary.star].ProtectSuccessNum}} times is sure to succeed
              </span>
              <span v-if="upAddData.upstarFailedNum != baseEquipStarData[ordinary.equip_type + '_' + ordinary.star].ProtectSuccessNum" :class="baseEquipStarData[ordinary.equip_type + '_' + ordinary.star].UpgradeSuccessRate < 10000 ? '' : 'success'">
                Success rate: {{baseEquipStarData[ordinary.equip_type + "_" + ordinary.star].UpgradeSuccessRate/100}}%
              </span>
              <span v-else class="success">
                Success rate: 100%
              </span>
            </div>
          </div>
        </div>
        <div class="foots">
          <div class="btxt hjicon" v-if="isUpgrade === 1">
            <p class="text-stroke">COST:&nbsp;<img src="../../assets/boss/goup.png"/><span>{{ materials.item10 }} / <span :class="[ hasEnoughMaterial ? 'enough' : 'notenough' ]">{{ costMaterial }}</span></span></p>
          </div>
          <div class="btxt hjicon" v-if="isUpgrade === 2 && sbxIndex === 1">
            <p class="text-stroke">COST:&nbsp;<img src="../../assets/items/11.png"/><span>{{ materials.item11 }} / <span :class="[ hasEnoughMaterial ? 'enough' : 'notenough' ]">{{ costMaterial }}</span></span></p>
          </div>
          <div class="btxt hjicon" v-if="isUpgrade === 2 && sbxIndex === 2">
            <p class="text-stroke">COST:&nbsp;<img src="../../assets/items/12.png"/><span>{{ materials.item12 }} / <span :class="[ hasEnoughMaterial ? 'enough' : 'notenough' ]">{{ costMaterial }}</span></span></p>
          </div>
          <div class="btxt hjicon" v-if="isUpgrade === 2 && sbxIndex === 3">
            <p class="text-stroke">COST:&nbsp;<img src="../../assets/items/13.png"/><span>{{ materials.item13 }} / <span :class="[ hasEnoughMaterial ? 'enough' : 'notenough' ]">{{ costMaterial }}</span></span></p>
          </div>
          <div class="btxt hjicon" v-if="isUpgrade === 2 && sbxIndex === 4">
            <p class="text-stroke">COST:&nbsp;<img src="../../assets/items/14.png"/><span>{{ materials.item14 }} / <span :class="[ hasEnoughMaterial ? 'enough' : 'notenough' ]">{{ costMaterial }}</span></span></p>
          </div>
          <div class="btxt hjicon" v-if="isUpgrade === 2 && sbxIndex === 5">
            <p class="text-stroke">COST:&nbsp;<img src="../../assets/items/15.png"/><span>{{ materials.item15 }} / <span :class="[ hasEnoughMaterial ? 'enough' : 'notenough' ]">{{ costMaterial }}</span></span></p>
          </div>
          <div v-if="!upgrading && hasEnoughMaterial" class="btnss f30" @click="onSubUpgrade(isUpgrade)">
            <div class="btn-redpoint" ></div>
            <span class="text-stroke">{{ isUpgrade === 1 ? 'Upgrade' : 'Enhance' }}</span>
          </div>
          <div v-if="!upgrading && !hasEnoughMaterial" class="btnss-unable f30">
            <span class="text-stroke">Not enough</span>
          </div>
          <div v-if="upgrading" class="upgrading">
            <div class="progress">
              <div class="doing" :style="{ width: upgradingVal + '%' }"></div>
            </div>
          </div>
        </div>
        <div class="close-btn" @click="onPopClose()"></div>
      </div>

      <div class="buffbody" v-if="isType === 5">
        <div v-if="this.isGettingBuff" class="getting text-stroke">
          <span class="dot"></span>
        </div>
        <div v-else v-for="(buff,index) in this.thisChoseBuff" :key="index" class="buff" :class="buffLevel(buff.level, choseAni,buff.id)" @click="choseBuff(buff.id)">
          <div class="bfimg">
            <img :src="buff.img" />
          </div>
          <div class="bfdesc">
            <div class="desc text-stroke">{{buff.desc}}</div>
          </div>
          <div v-if="!buff.round || buff.round == 0" class="bftime text-stroke">Use immediately</div>
          <div v-if="buff.round > 0" class="bftime text-stroke">Lasts {{buff.round}} fight{{buff.round === 1 ? '':'s'}}</div>
        </div>
      </div>

      <div class="popbody" v-if="isType === 6">
        <div class="head">
          <span class="text-stroke">My Current BUFF</span>
        </div>
        <div class="middle">
          <div class="c-buff" v-for="(buff,index) in this.buffCards" :key="index">
            <div class="cb-icon">
              <img :src="buff.img" />
            </div>
            <div class="cb-desc">
              <div class="desctxt text-stroke">{{buff.desc}}</div>
              <div class="buffleft text-stroke">
                Remaining:&nbsp;{{buff.leftround}}&nbsp;fight{{buff.leftround > 1 ? 's':''}}
              </div>
            </div>
          </div>
        </div>
        <div class="foots">
          <div class="btnss f30" style="margin-top: 25px;" @click="onPopClose()">
            <span class="text-stroke">Got it</span>
          </div>
        </div>
        <div class="close-btn" @click="onPopClose()"></div>
      </div>

      <!-- type 7 boss info -->
      <div class="popbody" v-if="isType === 7">
        <div class="head">
          <span class="text-stroke">BOSS Info</span>
        </div>
        <div class="middle">
          <div class="bs-detail">
            <div class="bs-pic">
              <img :src="this.thisBossData.picPath" />
            </div>
            <div class="bs-info">
              <div class="bs-name text-stroke">
                {{this.thisBossData.nameStr}}
                <span>Lv.{{this.thisBossData.Level}}</span>
              </div>
              <div class="bs-atr text-stroke">HP:&nbsp;{{this.thisBossData.BaseHP}}</div>
              <div class="bs-atr text-stroke">ATK:&nbsp;{{this.thisBossData.BaseAtk}}</div>
              <div class="bs-atr text-stroke">HIT:&nbsp;{{this.thisBossData.HitRate / 100}}%</div>
              <div class="bs-atr text-stroke">CRI:&nbsp;{{this.thisBossData.CriRate / 100}}%</div>
              <div class="bs-atr text-stroke">C.DMG:&nbsp;{{this.thisBossData.CriDmgRate / 100}}%</div>
            </div>
          </div>
        </div>
        <div class="foots">
          <div class="btnss f30" style="margin-top: 25px;" @click="onPopClose()">
            <span class="text-stroke">Close</span>
          </div>
        </div>
        <div class="close-btn" @click="onPopClose()"></div>
      </div>

      <div class="popbody" v-if="isType === 8">
        <div class="head">
          <span class="text-stroke">My Info</span>
        </div>
        <div class="middle">
          <div class="mybtl-base">
            <div class="avatar">
              <div class="pic"><img alt="avatar" src="~@/assets/character.png" /></div>
            </div>
            <div class="mybtl-atr">
              <div class="mybtl-name text-stroke">{{this.nickname}}</div>
              <div class="dataline text-stroke">HP:&nbsp;{{ this.battleinfo.hp }}<span>+{{ this.battleinfo.hpAdd }}</span></div>
              <div class="dataline text-stroke">ATK:&nbsp;{{ this.battleinfo.atk }}<span>+{{ this.battleinfo.atkAdd }}</span></div>
              <div class="dataline text-stroke">
                HIT:&nbsp;{{ Math.round(this.battleinfo.hit_rate / 100) }}%
              </div>
              <div class="dataline text-stroke">
                SPD:&nbsp;{{ Math.round(this.battleinfo.dodge_rate / 100) }}%
              </div>
              <div class="dataline text-stroke">
                CRT:&nbsp;{{ Math.round(this.battleinfo.cri_rate / 100) }}%
              </div>
              <div class="dataline text-stroke">
                C.DMG:&nbsp;{{ Math.round(this.battleinfo.cri_dmg_rate / 100) }}%
              </div>
            </div>
          </div>

          <div class="mybtl-equips">
            <div v-if="sbxData && sbxData.length > 0" class="base">
              <div class="baseequip" v-for="baseItem in sbxData" :key="baseItem.equip_type">
                <img v-if="baseItem.equip_type == 1 && baseItem.level < 3" alt="ori" src="~@/assets/boss/ordinary1.png" />
                <img v-if="baseItem.equip_type == 1 && baseItem.level >= 3 && baseItem.level < 9" alt="ori" src="~@/assets/boss/ordinary11.png" />
                <img v-if="baseItem.equip_type == 1 && baseItem.level >= 9" alt="ori" src="~@/assets/boss/ordinary12.png" />

                <img v-if="baseItem.equip_type == 2 && baseItem.level < 3" alt="ori" src="~@/assets/boss/ordinary2.png" />
                <img v-if="baseItem.equip_type == 2 && baseItem.level >= 3 && baseItem.level < 9" alt="ori" src="~@/assets/boss/ordinary21.png" />
                <img v-if="baseItem.equip_type == 2 && baseItem.level >= 9" alt="ori" src="~@/assets/boss/ordinary22.png" />

                <img v-if="baseItem.equip_type == 3 && baseItem.level < 3" alt="ori" src="~@/assets/boss/ordinary3.png" />
                <img v-if="baseItem.equip_type == 3 && baseItem.level >= 3 && baseItem.level < 9" alt="ori" src="~@/assets/boss/ordinary31.png" />
                <img v-if="baseItem.equip_type == 3 && baseItem.level >= 9" alt="ori" src="~@/assets/boss/ordinary32.png" />

                <img v-if="baseItem.equip_type == 4 && baseItem.level < 3" alt="ori" src="~@/assets/boss/ordinary4.png" />
                <img v-if="baseItem.equip_type == 4 && baseItem.level >= 3 && baseItem.level < 9" alt="ori" src="~@/assets/boss/ordinary41.png" />
                <img v-if="baseItem.equip_type == 4 && baseItem.level >= 9" alt="ori" src="~@/assets/boss/ordinary42.png" />

                <img v-if="baseItem.equip_type == 5 && baseItem.level < 3" alt="ori" src="~@/assets/boss/ordinary5.png" />
                <img v-if="baseItem.equip_type == 5 && baseItem.level >= 3 && baseItem.level < 9" alt="ori" src="~@/assets/boss/ordinary51.png" />
                <img v-if="baseItem.equip_type == 5 && baseItem.level >= 9" alt="ori" src="~@/assets/boss/ordinary52.png" />

                <div class="star"><span class="text-stroke">{{baseItem.star}}</span></div>
                <div class="lv"><span class="text-stroke">Lv.{{baseItem.level}}</span></div>
              </div>
            </div>
            <div v-else class="base">
              <span class="noinfo text-stroke">No base equipment found.</span>
            </div>
            <div v-if="this.nftEquipedArr && this.nftEquipedArr.length > 0" class="nfts">
              <div class="nftequip" v-for="nft in this.nftEquipedArr" :key="nft.id">
                <img alt="NFT" v-if="nft.level == 1 && nft.type == 1" src="~@/assets/nfts/Garm1.png" />
                <img alt="NFT" v-if="nft.level == 1 && nft.type == 2" src="~@/assets/nfts/Edda1.png" />
                <img alt="NFT" v-if="nft.level == 1 && nft.type == 3" src="~@/assets/nfts/Norns1.png" />
                <img alt="NFT" v-if="nft.level == 1 && nft.type == 4" src="~@/assets/nfts/Parvati1.png" />
                <img alt="NFT" v-if="nft.level == 1 && nft.type == 5" src="~@/assets/nfts/Tartarus1.png" />

                <img v-if="nft.level > 1 && nft.level <= 3 && nft.type == 1" alt="NFT" src="~@/assets/nfts/Garm2.png" />
                <img v-if="nft.level > 1 && nft.level <= 3 && nft.type == 2" alt="NFT" src="~@/assets/nfts/Edda2.png" />
                <img v-if="nft.level > 1 && nft.level <= 3 && nft.type == 3" alt="NFT" src="~@/assets/nfts/Norns2.png" />
                <img v-if="nft.level > 1 && nft.level <= 3 && nft.type == 4" alt="NFT" src="~@/assets/nfts/Parvati2.png" />
                <img v-if="nft.level > 1 && nft.level <= 3 && nft.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus2.png" />

                <img v-if="nft.level > 3 && nft.level <= 6 && nft.type == 1" alt="NFT" src="~@/assets/nfts/Garm3.png" />
                <img v-if="nft.level > 3 && nft.level <= 6 && nft.type == 2" alt="NFT" src="~@/assets/nfts/Edda3.png" />
                <img v-if="nft.level > 3 && nft.level <= 6 && nft.type == 3" alt="NFT" src="~@/assets/nfts/Norns3.png" />
                <img v-if="nft.level > 3 && nft.level <= 6 && nft.type == 4" alt="NFT" src="~@/assets/nfts/Parvati3.png" />
                <img v-if="nft.level > 3 && nft.level <= 6 && nft.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus3.png" />

                <img v-if="nft.level > 6 && nft.level <= 9 && nft.type == 1" alt="NFT" src="~@/assets/nfts/Garm4.png" />
                <img v-if="nft.level > 6 && nft.level <= 9 && nft.type == 2" alt="NFT" src="~@/assets/nfts/Edda4.png" />
                <img v-if="nft.level > 6 && nft.level <= 9 && nft.type == 3" alt="NFT" src="~@/assets/nfts/Norns4.png" />
                <img v-if="nft.level > 6 && nft.level <= 9 && nft.type == 4" alt="NFT" src="~@/assets/nfts/Parvati4.png" />
                <img v-if="nft.level > 6 && nft.level <= 9 && nft.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus4.png" />

                <img v-if="nft.level > 9 && nft.level <= 12 && nft.type == 1" alt="NFT" src="~@/assets/nfts/Garm5.png" />
                <img v-if="nft.level > 9 && nft.level <= 12 && nft.type == 2" alt="NFT" src="~@/assets/nfts/Edda5.png" />
                <img v-if="nft.level > 9 && nft.level <= 12 && nft.type == 3" alt="NFT" src="~@/assets/nfts/Norns5.png" />
                <img v-if="nft.level > 9 && nft.level <= 12 && nft.type == 4" alt="NFT" src="~@/assets/nfts/Parvati5.png" />
                <img v-if="nft.level > 9 && nft.level <= 12 && nft.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus5.png" />

                <img v-if="nft.level > 12 && nft.type == 1" alt="NFT" src="~@/assets/nfts/Garm6.png" />
                <img v-if="nft.level > 12 && nft.type == 2" alt="NFT" src="~@/assets/nfts/Edda6.png" />
                <img v-if="nft.level > 12 && nft.type == 3" alt="NFT" src="~@/assets/nfts/Norns6.png" />
                <img v-if="nft.level > 12 && nft.type == 4" alt="NFT" src="~@/assets/nfts/Parvati6.png" />
                <img v-if="nft.level > 12 && nft.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus6.png" />

                <div class="nftlv"><span class="text-stroke">{{nft.level}}</span></div>
              </div>
            </div>
            <div v-else class="nfts">
              <span class="noinfo text-stroke">No NFT weapon found.</span>
            </div>
          </div>
        </div>
        <div class="foots">
          <div class="btnss f30" style="margin-top: 25px;" @click="onPopClose()">
            <span class="text-stroke">Close</span>
          </div>
        </div>
        <div class="close-btn" @click="onPopClose()"></div>
      </div>
    </div>
    <div class="tips-div">
      <Tip ref="tip"></Tip>
    </div>
  </div>
</template>
<script>

// import { mapActions } from 'vuex'
import boss from '@/share/boss.json'
import equipLevel from '@/share/baseEquipLevels.json'
import equipStar from '@/share/baseEquipStars.json'
import initAttr from '@/share/playerInitAttrs.json'
import nftsInit from '@/share/nftsTmp.json'
import buffCardConfig from '@/share/BuffCard.json'
import languageConfig from '@/share/Language.json'
import jskit from '@/utils/index'
import foots from '@/components/foots'
import headers from '@/components/header'
import cache from '@/utils/cache'
import { getItems, upgradeBaseEquipStar, upgradeBaseEquipLevel } from '@/api/items'
import { getRandomBuff, choseBuff } from '@/api/fight'

// npm i @tonconnect/sdk
// https://ton-connect.github.io/sdk/modules/_tonconnect_sdk.html

// nRF Connect SDK-> https://docs.ton.org/develop/dapps/ton-connect/tg-bot-integration

export default {
  data() {
    const _this = this
    return {
      nickname:'',
      bossAni:0,
      bossChange: 0,
      isLastLevel: false,
      firstIntoPage: true,
      costMaterial: 0,
      hasEnoughMaterial: false,
      upgrading: false,
      upgradingVal: 0,
      upgradingFinish: false,
      upgradingSuccess: false,
      showcheck: false,
      initAttr: {},
      // 映射服务器attr结构，并增加自定义字段
      battleinfo: {
        hp: 0,
        atk: 0,
        hit_rate: 0,
        dodge_rate: 0,
        cri_rate: 0,
        cri_dmg_rate: 0,
        hp_add_rate: 0,
        atk_add_rate: 0,
        atkAdd: 0,
        hpAdd: 0
      },
      challenge: false,
      nftBossChallenge: {
        equipSameNft: false,
        upgradeOK: false,
        firstBoss: false
      },
      ons: 1,
      istip: false,
      isType: 0,
      // 材料数据
      materials: {
        item10: 0,
        item11: 0,
        item12: 0,
        item13: 0,
        item14: 0,
        item15: 0
      },
      // 模拟取数
      sbxIndex: 1,
      ordinary: {},
      isUpgrade: null,
      sbxUpData: [],
      sbxData: [
        {
          lv: 1,
          star: 1,
        },
        {
          lv: 1,
          star: 1,
        },
        {
          lv: 1,
          star: 1,
        },
        {
          lv: 1,
          star: 1,
        },
        {
          lv: 1,
          star: 1,
        }
      ],

      // nft equip
      hasChoicedNFT: false,
      nftEquip1: 0,
      nftEquip2: 0,
      nftEquip3: 0,
      nfts: {},
      nftEquipedArr:{},
      nftsListed:{},
      nftsEquiped1:{},
      nftsEquiped2:{},
      nftsEquiped3:{},
      nftChoiced:{},
      curNftIndex: 1,
      nftLevelTotal: 0,

      initBossLevel:{},
      bossData: {},
      thisClass: 0,
      classBossData: {},//同类型boss
      prevBossId: 0,
      thisBossId: 0,
      thisBossData: {},
      nextBossId: 0,
      firstRewardsData: {},
      initBossId: 0,
      baseEquipLevelData:{},
      baseEquipStarData:{},
      upAddData:{
        baseLvHpAdd: 0,
        baseLvAtkAdd: 0,
        baseStarHpAdd: 0,
        baseStarAtkAdd: 0,
        uplvFailedNum: 0,
        upstarFailedNum: 0
      },
      redpoint:{
        equipLv1: false,
        equipLv2: false,
        equipLv3: false,
        equipLv4: false,
        equipLv5: false,
        equipStar1: false,
        equipStar2: false,
        equipStar3: false,
        equipStar4: false,
        equipStar5: false
      },
      showupani: false,
      showuptxtani: false,
      showuptxtani2: false,
      showuptxtani3: false,

      buffStatus:0,
      buffCards:{},
      thisChoseBuff:{},
      isGettingBuff: false,
      choseAni: false,
      thisChosedId: 0
    }
  },
  components: { foots, headers },
  computed: {
    swipeMode() {
      return function(index){
        if(index === 1){
          return 'slideout-next'
        }else if(index === 2){
          return 'slidein-next'
        }else if(index === 3){
          return 'slideout-prev'
        }else if(index === 4){
          return 'slidein-prev'
        }else{
          return ''
        }
      }
    },
    bossChangeMode(){
      return function(index){
        if(index === 1){
          return 'change-boss'
        }else{
          return ''
        }
      }
    },
    buffLevel(){
      return function(level,choseAni,id){
        let cls = ''
        if(level === 1){
          cls = 'blue'
        }else if(level === 2){
          cls = 'purple'
        }else if(level === 3){
          cls = 'yellow'
        }
        if(choseAni){
          cls += id == this.thisChosedId ? ' choseAni' : ' leaveAni'
        }
        return cls
      }
    }

  },
  mounted() {
    const querys = this.$route.query || {}
    this.firstIntoPage = true
    //boss数据
    this.bossData = boss

    this.nickname = cache.getSession('nickName') || 'Player'

    // 读表装备数据
    this.baseEquipLevelData = equipLevel
    this.baseEquipStarData = equipStar
  },
  activated(){
    // 每次进入激活
    console.log("active")
    // 初始化NFT装备
    this.nfts = jskit.localStorageGet('nftsTmp')
    if(!this.nfts){
      this.nfts = nftsInit
      jskit.localStorageSet('nftsTmp', this.nfts)
    }
    this.initBossLevelInfo()
    this.initBuffCard()
    this.initNftEquip()
    this.freshNftLevelTotal()
    this.getSbxData()
    this.freshBoss()

  },
  methods: {
    onPopBox(type) {
      console.log(type, '====type')
      this.istip = true
      this.isType = parseInt(type)
    },
    initBossLevelInfo(){
      this.initBossLevel = {}
      console.log("boss:", boss)
      const map = new Map(Object.entries(boss));
      console.log("boss total num:",map.size)
      let levelArr = []
      for (let [key, value] of map) {
        // 大于20000是隐藏BOSS  1-1 第一关
        let _boss = {}
        if(key < 20000 && value.LevelShow == '1-1'){
          _boss.class = value.Class
          _boss.startId = value.BossId
          levelArr.push(_boss)
        }
      }
      for (let [key, value] of map) {
        // 大于20000是隐藏BOSS  15-3 最后一关
        if(key < 20000 && value.LevelShow == '15-3'){
          levelArr.find(obj => obj.class === value.Class).endId = value.BossId
        }
      }
      this.initBossLevel = levelArr
      // console.log(this.initBossLevel)
    },
    initBuffCard(){
      this.buffStatus =+ cache.getSession('buffCardStatus') || 0
      this.thisChoseBuff = {}
      this.isGettingBuff = false
      this.buffCards = {}
      // 刷新左上角BUFF
      let cards = JSON.parse(cache.getSession('buffCard')) || {}
      if(cards && cards.length > 0){
        let buffArrToChose = []
        for(let bf of cards){
          let buff = {}
          buff.id = buffCardConfig[bf.buff_card_id].CardId
          buff.type = buffCardConfig[bf.buff_card_id].CardType
          buff.level = buffCardConfig[bf.buff_card_id].CardLevel
          buff.leftround = bf.left_round
          buff.img = require('@/assets/buff/' + buff.id + '.png')
          buff.desc = languageConfig[buffCardConfig[bf.buff_card_id].Desc].EN || ''
          buffArrToChose.push(buff)
        }
        this.buffCards = buffArrToChose
      }
    },
    showBuffCards(bf_cards){
      let buffArrToChose = []
      for(let bf of bf_cards){
        let buff = {}
        buff.id = buffCardConfig[bf].CardId
        buff.type = buffCardConfig[bf].CardType
        buff.level = buffCardConfig[bf].CardLevel
        buff.round = buffCardConfig[bf].Round
        buff.img = require('@/assets/buff/' + buff.id + '.png')
        buff.desc = languageConfig[buffCardConfig[bf].Desc].EN || ''
        buffArrToChose.push(buff)
      }
      this.thisChoseBuff = buffArrToChose
      this.isGettingBuff = false
      this.choseAni = false
      this.thisChosedId = 0
    },
    onStartBtn() {
      // location.href = `/pk?id=${this.thisBossId}`
      let energy =+ cache.getSession('energy') || 0
      if(this.thisBossData.EnergyCost > energy){
        this.$refs.tip.openTip('Energy not enough')
      }else{
        // 判断BUFF情况
        if(this.buffStatus === 0 || this.buffStatus === 1){
          let bid =+ this.thisBossId
          this.thisChoseBuff = {}
          this.choseAni = false
          this.thisChosedId = 0
          this.isGettingBuff = true
          this.onPopBox(5)
          console.log("===boss id ", bid)
          getRandomBuff({boss_id : bid}).then(response => {
            console.log("fresh buff ", bid, response, response.success)
            if(response && response.buff_cards && response.buff_cards.length === 3){
              this.showBuffCards(response.buff_cards)
            }else{
              this.$refs.tip.openTip('Get Buff Error')
            }
          })
        }else if(this.buffStatus === 2){
          this.$router.push({ path: '/pk', query: { 'id': this.thisBossId || '', 'name':  this.thisBossData.nameStr} })
        }
      }
    },
    choseBuff(buffId){
      let id =+ buffId
      this.choseAni = true
      this.thisChosedId = buffId
      choseBuff({buff_card : id}).then(response => {
        console.log("chose", response)
        if(response && response.buff_cards && response.buff_cards.length === 3){
          // 重新刷三张卡
          this.thisChoseBuff = {}
          this.showBuffCards(response.buff_cards)
        }else{
          // 开始打BOSS
          this.onPopClose()
          this.$router.push({ path: '/pk', query: { 'id': this.thisBossId || '', 'name':  this.thisBossData.nameStr} })
        }
      })
    },
    buttonPrev() {
      //给新选的boss赋值
      console.log("+++",this.thisBossData)
      if (this.prevBossId) {
        this.bossAni = 3
        let timer = setTimeout(() => {
          this.getThisBossData(this.prevBossId)
          this.bossAni = 4
          let timerFinish = setTimeout(() => {
            this.bossAni = 0
            clearTimeout(timerFinish)
          }, 400)
          clearTimeout(timer)
        }, 200)
      }
    },
    buttonNext() {
      //给新选的boss赋值
      console.log("+++",this.thisBossData)
      if (this.nextBossId) {
        this.bossAni = 1
        let timer = setTimeout(() => {
          this.getThisBossData(this.nextBossId)
          this.bossAni = 2
          let timerFinish = setTimeout(() => {
            this.bossAni = 0
            clearTimeout(timerFinish)
          }, 400)
          clearTimeout(timer)
        }, 200)
      }
    },
    //选择一个类型boss
    buttonOne(theIndex) {
      if (theIndex > 0 && theIndex < 7) {
        //不点当前菜单
        if (theIndex - 1 != this.thisClass) {
          //初始化boss信息
          this.ons = theIndex
          this.thisClass = this.ons - 1
          // 进入选中的类型中，最新可以挑战的BOSS等级
          this.calNextNode(this.thisClass)
          this.getClassBossData()
          this.bossChange = 1
          let timer = setTimeout(() => {
            this.bossChange = 0
            clearTimeout(timer)
          }, 300)
        }
      }
    },
    calNextNode(classType){
      this.isLastLevel = false
      let nodes = JSON.parse(cache.getSession('bossNode')) || {}
      let _node = nodes.find(obj => obj.class === classType).bossId
      let startId = this.initBossLevel.find(obj => obj.class === classType).startId
      let endId = this.initBossLevel.find(obj => obj.class === classType).endId
      if(_node == endId){
        this.isLastLevel = true
      }else{
        if(_node == 0){
          _node = startId
        }else{
          _node += 1
        }
      }
      this.thisBossId =+ _node
      this.initBossId =+ _node
    },
    freshBoss(){
      const querys = this.$route.query || {}
      if(querys.id){
        this.firstIntoPage = false
        this.thisBossId = +querys.id
        this.initBossId = +querys.id
        if (this.thisBossId) {
          this.thisClass = this.bossData[this.thisBossId] ? this.bossData[this.thisBossId].Class : 0
        }
        console.log("from fight, bossid=",this.initBossId," class=",this.thisClass)
        if(querys.win && querys.win == 1){
          // 胜利时往后进一位
          this.calNextNode(this.thisClass)
        }
      }else{
        console.log(this.bossId,'====this bossid，class===',this.thisClass)
        this.calNextNode(this.thisClass)
      }
      if (this.thisClass + 1 != this.ons) {
        // 刷新右侧tab图标
        this.ons = this.thisClass + 1
      }
      console.log(this.firstIntoPage)
      if (!this.firstIntoPage && !this.isLastLevel && querys.win && querys.win === 1) {
        // 播动画
        this.bossAni = 1
        let timer = setTimeout(() => {
          this.getClassBossData()
          this.bossAni = 2
          let timerFinish = setTimeout(() => {
            this.bossAni = 0
            clearTimeout(timerFinish)
          }, 400)
          clearTimeout(timer)
        }, 200)
      }else{
        this.getClassBossData()
      }
    },
    freshMaterial(bagitems) {
      if (bagitems && bagitems.length > 0) {
        for (let item of bagitems) {
          if (item.item_id == 10) {
            this.materials.item10 = item.num
          } else if (item.item_id == 11) {
            this.materials.item11 = item.num
          } else if (item.item_id == 12) {
            this.materials.item12 = item.num
          } else if (item.item_id == 13) {
            this.materials.item13 = item.num
          } else if (item.item_id == 14) {
            this.materials.item14 = item.num
          } else if (item.item_id == 15) {
            this.materials.item15 = item.num
          }
        }
      }
    },
    getSbxData() {
      let _sbxData = JSON.parse(cache.getSession('baseEquip')) || {}
      let _sbxUpData = JSON.parse(cache.getSession('baseUpInfo')) || {}
      let _materials = JSON.parse(cache.getSession('bagItems')) || {}
      // 优先从session中取值
      if(Object.keys(_sbxData).length != 0 && Object.keys(_sbxUpData).length != 0 && Object.keys(_materials).length != 0){
        console.log("get items from session")
        this.sbxData = _sbxData
        this.sbxUpData = _sbxUpData
        this.freshMaterial(_materials)
        this.freshRedpoint()
        this.calBattleInfo()
      }else{
        getItems({}).then(response => {
          console.log("get items from api", response)
          this.sbxData = response.base_equips
          this.sbxUpData = response.base_equip_upgrade_infos
          this.freshMaterial(response.items)
          console.log(this.materials)
          if(response.base_equips){
            cache.setSession('baseEquip', response.base_equips)
          }
          if(response.base_equip_upgrade_infos){
            cache.setSession('baseUpInfo', response.base_equip_upgrade_infos)
          }
          if(response.items){
            cache.setSession('bagItems', response.items)
          }
          if(response.mana){
            cache.setSession('zpa', response.mana)
            this.$store.commit("GLOBAL_ZPA", response.mana)
          }
          this.freshRedpoint()
          this.calBattleInfo()
        })
      }
    },
    //获取这个级别boss所有boss数据
    getClassBossData() {
      //初始化级别
      this.classBossData = {}
      for (let bossId in this.bossData) {
        if (this.bossData[bossId].Class == this.thisClass && bossId < 20000) {
          // console.log('class++++', this.bossData[bossId].Class, this.thisClass)
          this.classBossData[bossId] = this.bossData[bossId]
        }
      }
      // 初始化boss信息
      this.getThisBossData(0)
      // console.log('class=======this.thisBossId', this.thisBossId)
      // console.log('class======this.classBossData', this.classBossData)
    },
    //获取当前boss信息
    getThisBossData(bossId) {
      //初始化
      this.firstRewardsData = {}
      this.prevBossId = 0
      this.thisBossId = 0
      this.nextBossId = 0
      let classBossIds = Object.keys(this.classBossData)
      if (!bossId) {
        if (!this.initBossId) {
          //新进游戏，没有初始值
          this.thisBossId = classBossIds[0]
          this.thisBossData = this.classBossData[this.thisBossId]
          this.nextBossId = classBossIds[1]
        } else {
          //打完boss有初始值
          for (let k in classBossIds) {
            //找到这个boss,上一个和下一个
            if (classBossIds[+k] == this.initBossId) {
              this.thisBossId = classBossIds[+k]
              this.thisBossData = this.classBossData[this.initBossId]
              this.prevBossId = classBossIds[+k - 1] ? classBossIds[+k - 1] : 0
              this.nextBossId = classBossIds[+k + 1] ? classBossIds[+k + 1] : 0
            }
          }
          this.initBossId = 0
        }
      } else {
        // 上一个与下一个按键的处理逻辑
        for (let k in classBossIds) {
          //找到这个boss,上一个和下一个
          if (classBossIds[+k] == bossId) {
            this.thisBossId = classBossIds[+k]
            this.thisBossData = this.classBossData[this.thisBossId]
            this.prevBossId = classBossIds[+k - 1] ? classBossIds[+k - 1] : 0
            this.nextBossId = classBossIds[+k + 1] ? classBossIds[+k + 1] : 0
          }
        }
      }
      //获取能挑战boss
      this.canChallenge(this.thisBossId, this.thisClass)
      this.canChallengeNftBoss()

      //首通奖励
      for (let itemId in this.thisBossData.FirstWinAwards) {
        if (!this.firstRewardsData[itemId]) {
          this.firstRewardsData[itemId] = {}
        }
        this.firstRewardsData[itemId]['itemId'] = itemId
        this.firstRewardsData[itemId]['num'] = this.thisBossData.FirstWinAwards[itemId]
        this.firstRewardsData[itemId]['itemPic'] = require('@/assets/items/' + itemId + '.png')
      }
      //普通奖励
      for (let itemId in this.thisBossData.Awards) {
        this.thisBossData.Awards[itemId]['itemPic'] = require('@/assets/items/' + itemId + '.png')
      }
      if(this.thisBossData.ManaAwards){
        this.thisBossData.ManaAwardsFinal = 0
        let mana = this.thisBossData.ManaAwards.split(":")
        if(mana.length == 2){
          this.thisBossData.ManaAwardsFinal = mana[0]
        }
      }
      this.thisBossData.picPath = require('@/assets/boss/' + this.thisBossData.Image + '.png')
      this.thisBossData.picShadowPath = require('@/assets/boss/' + this.thisBossData.Image + 'b.png')
      this.thisBossData.nameStr = languageConfig[this.thisBossData.Name].EN || ''

      console.log(this.thisBossData,"+++=====")
      // console.log('this.prevBossId', this.prevBossId)
      // console.log('this.thisBossId', this.thisBossId)
      // console.log('this.nextBossId', this.nextBossId)
    },
    canChallenge(bossId, bossClass) {
      if (bossId != 0) {
        let nodes = JSON.parse(cache.getSession('bossNode')) || {}
        let _node = nodes.find(obj => obj.class === bossClass).bossId
        let _initId = this.initBossLevel.find(obj => obj.class === bossClass).startId

        if (_node == 0) { // 没有进度
          this.showcheck = false;
          if (bossId == _initId) {
            this.challenge = true;
          } else {
            this.challenge = false;
          }
        } else {
          this.showcheck = false;
          if (bossId - 1 <= _node) {
            this.challenge = true;
            if (bossId <= _node) {
              this.showcheck = true;
            }
          } else {
            this.showcheck = false;
            this.challenge = false;
          }
        }
      }
    },
    canChallengeNftBoss(){
      this.nftBossChallenge.equipSameNft = false
      this.nftBossChallenge.upgradeOK = false
      this.nftBossChallenge.firstBoss = false
      if(this.thisBossId != 0){
        if(this.thisClass == 0){
          // class为0时为普通BOSS，无需处理
          this.nftBossChallenge.equipSameNft = true
          this.nftBossChallenge.upgradeOK = true
        }else{
          let _boss = this.classBossData[this.thisBossId]
          let nftArr = []
          if(this.nftsEquiped1){
            nftArr.push(this.nftsEquiped1)
          }
          if(this.nftsEquiped2){
            nftArr.push(this.nftsEquiped2)
          }
          if(this.nftsEquiped3){
            nftArr.push(this.nftsEquiped3)
          }
          console.log(this.nftsEquiped1,this.nftsEquiped2,this.nftsEquiped3)
          console.log("current boss equip level need", _boss.NFTEquipsLevel,"equip nfts", nftArr,"boss class", this.thisClass)
          if(nftArr && nftArr.length > 0){
            for(let nft of nftArr){
              console.log(nft.type)
              // 由于NFT是临时数据，nft type与boss class没有做一一对应，后续需优化
              if(this.thisClass == 1 && nft.type == 1){
                this.nftBossChallenge.equipSameNft = true
              }
              if(this.thisClass == 2 && nft.type == 3){
                this.nftBossChallenge.equipSameNft = true
              }
              if(this.thisClass == 3 && nft.type == 5){
                this.nftBossChallenge.equipSameNft = true
              }
              if(this.thisClass == 4 && nft.type == 4){
                this.nftBossChallenge.equipSameNft = true
              }
              if(this.thisClass == 5 && nft.type == 2){
                this.nftBossChallenge.equipSameNft = true
              }
            }
            if(this.nftLevelTotal >= _boss.NFTEquipsLevel){
              this.nftBossChallenge.upgradeOK = true
            }
          }
          if(this.thisBossId == '11001' || this.thisBossId == '12001' || this.thisBossId == '13001'
            || this.thisBossId == '14001' || this.thisBossId == '15001'){
            this.nftBossChallenge.firstBoss = true
          }
        }
      }
    },
    initAddData(){
      this.upAddData.baseLvHpAdd = 0
      this.upAddData.baseLvAtkAdd = 0
      this.upAddData.baseStarHpAdd = 0
      this.upAddData.baseStarAtkAdd = 0
      this.upAddData.uplvFailedNum = 0
      this.upAddData.upstarFailedNum = 0
    },
    calUpData(index){
      // 计算升级升星增加的HP与ATK
      this.initAddData()
      if(this.ordinary){
        if(this.ordinary.level < 20){
          let curLv = this.baseEquipLevelData[this.ordinary.equip_type + "_" + this.ordinary.level]
          let nextLv = this.baseEquipLevelData[this.ordinary.equip_type + "_" + (this.ordinary.level + 1)]
          this.upAddData.baseLvHpAdd = nextLv.BaseHP - curLv.BaseHP
          this.upAddData.baseLvAtkAdd = nextLv.BaseAtk - curLv.BaseAtk
          if(this.sbxUpData[+index - 1] && this.sbxUpData[+index - 1].equip_type == index){
            this.upAddData.uplvFailedNum = this.sbxUpData[+index - 1].upgrade_level_failed_times || 0
          }
        }
        if(this.ordinary.star < 20){
          let curStar = this.baseEquipStarData[this.ordinary.equip_type + "_" + this.ordinary.star]
          let nextStar = this.baseEquipStarData[this.ordinary.equip_type + "_" + (this.ordinary.star + 1)]
          this.upAddData.baseStarHpAdd = nextStar.BaseHP - curStar.BaseHP
          this.upAddData.baseStarAtkAdd = nextStar.BaseAtk - curStar.BaseAtk
          if(this.sbxUpData[+index - 1] && this.sbxUpData[+index - 1].equip_type == index){
            this.upAddData.upstarFailedNum = this.sbxUpData[+index - 1].upgrade_star_failed_times || 0
          }
        }
      }
    },
    //打开装备详情
    onSbx(index) {
      this.showupani = false
      this.showuptxtani = false
      this.upgrading = false
      this.onPopBox(3)
      this.sbxIndex = index
      this.getSbxData()
      //init data ,star lv hp atk
      this.ordinary = this.sbxData[+index - 1]
      console.log(this.ordinary, 'this.ordinary')
      this.calUpData(index);
    },
    // 刷新消耗
    freshCost(type) {
      this.hasEnoughMaterial = false;
      if (this.ordinary) {
        if (type == 1) {
          let curLv = this.baseEquipLevelData[this.ordinary.equip_type + "_" + this.ordinary.level]
          this.costMaterial = curLv.UpgradeConsumeItems[10]
          if (this.costMaterial <= this.materials.item10) {
            this.hasEnoughMaterial = true;
          }
        } else if (type == 2) {
          let curStar = this.baseEquipStarData[this.ordinary.equip_type + "_" + this.ordinary.star]
          if (this.sbxIndex == 1) {
            this.costMaterial = curStar.UpgradeConsumeItems[11]
            if (this.costMaterial <= this.materials.item11) {
              this.hasEnoughMaterial = true;
            }
          } else if (this.sbxIndex == 2) {
            this.costMaterial = curStar.UpgradeConsumeItems[12]
            if (this.costMaterial <= this.materials.item12) {
              this.hasEnoughMaterial = true;
            }
          } else if (this.sbxIndex == 3) {
            this.costMaterial = curStar.UpgradeConsumeItems[13]
            if (this.costMaterial <= this.materials.item13) {
              this.hasEnoughMaterial = true;
            }
          } else if (this.sbxIndex == 4) {
            this.costMaterial = curStar.UpgradeConsumeItems[14]
            if (this.costMaterial <= this.materials.item14) {
              this.hasEnoughMaterial = true;
            }
          } else if (this.sbxIndex == 5) {
            this.costMaterial = curStar.UpgradeConsumeItems[15]
            if (this.costMaterial <= this.materials.item15) {
              this.hasEnoughMaterial = true;
            }
          }
        }
      }
    },
    upgradeData(type) {
      // 从装备详情进入到升级或升星界面
      this.freshCost(type);
      this.istip = false;
      this.onPopBox(4)
      //1upgrade, 2enhance
      this.isUpgrade = type;
      //普通升级最高20级；
      //皇冠15级；
    },
    //普通装备升级升星
    onSubUpgrade(type) {
      //type 1 Lv++ ,2 stat++
      // console.log(this.ordinary, '----')
      // 升级开始加载进度条
      this.upgradingFinish = false
      this.upgradingSuccess = false
      this.showUpgradingProgress(type)
      let equipType = this.ordinary.equip_type
      if (type === 1) {
        let data = {
          equipt_type: equipType,
          cur_level: this.ordinary.level
        }
        upgradeBaseEquipLevel(data).then(response => {
          console.log('up response', response)
          // 无论成败，先更新材料剩余数量，刷新session
          this.materials.item10 = this.materials.item10 - this.costMaterial
          let _materials = JSON.parse(cache.getSession('bagItems')) || {}
          if(Object.keys(_materials).length != 0){
            _materials.find(obj => obj.item_id === 10).num = this.materials.item10;
            cache.setSession('bagItems', _materials)
          }
          if (response.success) {
            console.log('success', response)
            this.doUpSuccessResponse(type, equipType)
          } else {
            console.log('failed', response)
            this.doUpFailedResponse(type, equipType)
          }
        })
      }
      if (type === 2) {
        let data = {
          equipt_type: equipType,
          cur_star: this.ordinary.star
        }
        upgradeBaseEquipStar(data).then(response => {
          console.log('up star response', response)
          // 无论成败，先更新材料剩余数量，刷新session
          let idx = 0
          let _mat = 0
          if(equipType == 1){
            this.materials.item11 = this.materials.item11 - this.costMaterial
            idx = 11
            _mat = this.materials.item11
          }else if(equipType == 2){
            this.materials.item12 = this.materials.item12 - this.costMaterial
            idx = 12
            _mat = this.materials.item12
          }else if(equipType == 3){
            this.materials.item13 = this.materials.item13 - this.costMaterial
            idx = 13
            _mat = this.materials.item13
          }else if(equipType == 4){
            this.materials.item14 = this.materials.item14 - this.costMaterial
            idx = 14
            _mat = this.materials.item14
          }else if(equipType == 5){
            this.materials.item15 = this.materials.item15 - this.costMaterial
            idx = 15
            _mat = this.materials.item15
          }
          let _materials = JSON.parse(cache.getSession('bagItems')) || {}
          if(Object.keys(_materials).length != 0){
            _materials.find(obj => obj.item_id === idx).num = _mat
            cache.setSession('bagItems', _materials)
          }
          if (response.success) {
            console.log('success', response)
            this.doUpSuccessResponse(type, equipType)
          } else {
            console.log('failed', response)
            this.doUpFailedResponse(type, equipType)
          }
        })
      }
    },

    doUpSuccessResponse(type, equipType){
      // 先更新数值，再重新计算下一级增加的数组
      console.log("add Lv num ",this.upAddData.baseLvHpAdd,this.upAddData.baseLvAtkAdd)
      console.log("add star num ",this.upAddData.baseStarHpAdd,this.upAddData.baseStarAtkAdd)
      let _sbxData = JSON.parse(cache.getSession('baseEquip')) || {}
      if(Object.keys(_sbxData).length != 0){
        if(type === 1){
          _sbxData.find(obj => obj.equip_type == equipType).level += 1;
          _sbxData.find(obj => obj.equip_type == equipType).attrs.hp += this.upAddData.baseLvHpAdd;
          _sbxData.find(obj => obj.equip_type == equipType).attrs.atk += this.upAddData.baseLvAtkAdd;
        }else if(type === 2){
          _sbxData.find(obj => obj.equip_type == equipType).star += 1;
          _sbxData.find(obj => obj.equip_type == equipType).attrs.hp += this.upAddData.baseStarHpAdd;
          _sbxData.find(obj => obj.equip_type == equipType).attrs.atk += this.upAddData.baseStarAtkAdd;
        }
        this.sbxData = _sbxData
        cache.setSession('baseEquip', _sbxData)
      }
      let _sbxUpData = JSON.parse(cache.getSession('baseUpInfo')) || {}
      if(Object.keys(_sbxUpData).length != 0){
        if(type === 1){
          if(_sbxUpData.find(obj => obj.equip_type == equipType).upgrade_level_failed_times){
            _sbxUpData.find(obj => obj.equip_type == equipType).upgrade_level_failed_times = 0
          }
        }else if(type === 2){
          if(_sbxUpData.find(obj => obj.equip_type == equipType).upgrade_star_failed_times){
            _sbxUpData.find(obj => obj.equip_type == equipType).upgrade_star_failed_times = 0
          }
        }
        this.sbxUpData = _sbxUpData
        cache.setSession('baseUpInfo', _sbxUpData)
      }
      this.ordinary = this.sbxData[+equipType - 1]
      this.freshCost(type)
      this.freshRedpoint()
      this.calUpData(equipType)
      this.calBattleInfo()
      this.upgradingSuccess = true
      this.upgradingFinish = true
    },

    doUpFailedResponse(type, equipType){
      // 更新失败次数
      let _sbxUpData = JSON.parse(cache.getSession('baseUpInfo')) || {}
      if(Object.keys(_sbxUpData).length != 0){
        if(type === 1){
          if(_sbxUpData.find(obj => obj.equip_type == equipType).upgrade_level_failed_times){
            _sbxUpData.find(obj => obj.equip_type == equipType).upgrade_level_failed_times += 1
          }else{
            _sbxUpData.find(obj => obj.equip_type == equipType)['upgrade_level_failed_times'] = 1
          }
        }else if(type === 2){
          if(_sbxUpData.find(obj => obj.equip_type == equipType).upgrade_star_failed_times){
            _sbxUpData.find(obj => obj.equip_type == equipType).upgrade_star_failed_times += 1
          }else{
            _sbxUpData.find(obj => obj.equip_type == equipType)['upgrade_star_failed_times'] = 1
          }
        }
        this.sbxUpData = _sbxUpData
        cache.setSession('baseUpInfo', _sbxUpData)
      }
      this.freshCost(type)
      this.freshRedpoint()
      this.calUpData(equipType)
      this.upgradingFinish = true
    },

    // NFT 模块
    initNftEquip(){
      this.nftEquipedArr = {}
      let arr = []
      this.nftEquip1 = 0
      this.nftEquip2 = 0
      this.nftEquip3 = 0
      this.nftsEquiped1 = jskit.localStorageGet('equip1')
      this.nftsEquiped2 = jskit.localStorageGet('equip2')
      this.nftsEquiped3 = jskit.localStorageGet('equip3')
      if(this.nftsEquiped1){
        this.nftEquip1 = this.nftsEquiped1.id
        arr.push(this.nftsEquiped1)
      }
      if(this.nftsEquiped2){
        this.nftEquip2 = this.nftsEquiped2.id
        arr.push(this.nftsEquiped2)
      }
      if(this.nftsEquiped3){
        this.nftEquip3 = this.nftsEquiped3.id
        arr.push(this.nftsEquiped3)
      }
      this.nftEquipedArr = arr
    },
    freshNftLevelTotal(){
      this.nftEquipedArr = {}
      let arr = []
      this.nftLevelTotal = 0
      this.nftsEquiped1 = jskit.localStorageGet('equip1')
      this.nftsEquiped2 = jskit.localStorageGet('equip2')
      this.nftsEquiped3 = jskit.localStorageGet('equip3')
      if(this.nftsEquiped1){
        this.nftLevelTotal += this.nftsEquiped1.level
        arr.push(this.nftsEquiped1)
      }
      if(this.nftsEquiped2){
        this.nftLevelTotal += this.nftsEquiped2.level
        arr.push(this.nftsEquiped2)
      }
      if(this.nftsEquiped3){
        this.nftLevelTotal += this.nftsEquiped3.level
        arr.push(this.nftsEquiped3)
      }
      this.nftEquipedArr = arr
    },
    choiceNft(id){
      for(let nft of this.nfts.weapon){
        if(nft.id == id){
          this.nftChoiced = nft
        }
      }
      this.hasChoicedNFT = true
    },
    equipNft(){
      if(this.nftChoiced){
        if(this.curNftIndex == 1){
          this.nftEquip1 = this.nftChoiced.id
          this.nftsEquiped1 = this.nftChoiced
          jskit.localStorageSet('equip1', this.nftChoiced)
        }else if(this.curNftIndex == 2){
          this.nftEquip2 = this.nftChoiced.id
          this.nftsEquiped2 = this.nftChoiced
          jskit.localStorageSet('equip2', this.nftChoiced)
        }else if(this.curNftIndex == 3){
          this.nftEquip3 = this.nftChoiced.id
          this.nftsEquiped3 = this.nftChoiced
          jskit.localStorageSet('equip3', this.nftChoiced)
        }
        this.freshNftLevelTotal()
        this.canChallengeNftBoss()
        this.calBattleInfo()
        this.istip = false
        this.isType = 0
      }else{
        this.$refs.tip.openTip('Param Error')
      }
    },
    openNftEquip(index){
      this.curNftIndex = index
      this.hasChoicedNFT = false
      this.nftChoiced = {}
      if(index == 1 && this.nftEquip1 > 0){
        this.hasChoicedNFT = true
        this.choiceNft(this.nftEquip1)
      }
      if(index == 2 && this.nftEquip2 > 0){
        this.hasChoicedNFT = true
        this.choiceNft(this.nftEquip2)
      }
      if(index == 3 && this.nftEquip3 > 0){
        this.hasChoicedNFT = true
        this.choiceNft(this.nftEquip3)
      }
      if(this.nfts && this.nfts.weapon && this.nfts.weapon.length > 0){
        let _arr = []
        _arr.push(this.nftChoiced)
        for(let nft of this.nfts.weapon){
          if(nft.id != this.nftChoiced.id){
            _arr.push(nft)
          }
        }
        this.nftsListed = _arr
      }
      this.onPopBox(1)
    },
    onOpenNftUp(type) {
      this.istip = false
      this.showuptxtani = false
      this.onPopBox(2)
      //1upgrade, 2enhance
      this.isUpgrade = type
      //普通升级最高20级；
      //皇冠15级；
    },
    onUpgradeNFT() {
      this.showuptxtani = false
      this.showuptxtani2 = false
      this.showuptxtani3 = false
      //type 1 Lv++ ,2 stat++
      let zpa =+ cache.getSession('zpa') || 0
      if(zpa < 1){
        this.$refs.tip.openTip('Token is not enough')
      }else{
        if(this.nftChoiced.level >= 15){
          this.$refs.tip.openTip('Highest level')
        }else{
          zpa = zpa - 1
          cache.setSession('zpa', zpa)
          this.$store.commit("GLOBAL_ZPA", zpa)

          if(this.nftChoiced && this.nftChoiced.level < 15){
            // 基础属性+100 lv+1
            this.nftChoiced.level += 1;
            this.nftChoiced.main_atr_1_val += 100;
            this.nftChoiced.main_atr_2_val += 100;
            this.showuptxtani = true
            if(this.nftChoiced.level == 3 || this.nftChoiced.level == 6 ||
            this.nftChoiced.level == 9 || this.nftChoiced.level == 12 || this.nftChoiced.level == 15){
              console.log('add random attr')
              let attr = Math.floor((Math.random()*5)+1) * 100;
              let type = Math.floor((Math.random()*2)+1)
              // type = 1 ATK 2 HP
              let addAtr1 = this.nftChoiced.add_atr_1
              if(addAtr1 == 'none'){
                if(type == 1){
                  addAtr1 = 'ATK'
                }else if(type == 2){
                  addAtr1 = 'HP'
                }
                this.nftChoiced.add_atr_1 = addAtr1
                this.nftChoiced.add_atr_1_val = attr
                this.showuptxtani2 = true
              }else{
                if(type == 1){
                  if(addAtr1 == 'ATK'){
                    this.nftChoiced.add_atr_1_val += attr
                    this.showuptxtani2 = true
                  }else{
                    this.nftChoiced.add_atr_2 = 'ATK'
                    this.nftChoiced.add_atr_2_val += attr
                    this.showuptxtani3 = true
                  }
                }
                if(type == 2){
                  if(addAtr1 == 'HP'){
                    this.nftChoiced.add_atr_1_val += attr
                    this.showuptxtani2 = true
                  }else{
                    this.nftChoiced.add_atr_2 = 'HP'
                    this.nftChoiced.add_atr_2_val += attr
                    this.showuptxtani3 = true
                  }
                }
              }
            }
            // 更新 storage与缓存
            this.nfts = jskit.localStorageGet('nftsTmp')
            if(this.nfts && this.nfts.weapon){
              let idx = 0
              for(let nft of this.nfts.weapon){
                if(nft.id == this.nftChoiced.id){
                  this.nfts.weapon[idx] = this.nftChoiced
                }
                idx++
              }
              jskit.localStorageSet('nftsTmp', this.nfts)
            }
            if(this.nftChoiced.id == this.nftEquip1){
              this.nftsEquiped1 == this.nftChoiced
              jskit.localStorageSet('equip1', this.nftChoiced)
            }else if(this.nftChoiced.id == this.nftEquip2){
              this.nftsEquiped2 == this.nftChoiced
              jskit.localStorageSet('equip2', this.nftChoiced)
            }else if(this.nftChoiced.id == this.nftEquip3){
              this.nftsEquiped3 == this.nftChoiced
              jskit.localStorageSet('equip3', this.nftChoiced)
            }
            this.freshNftLevelTotal()
            this.canChallengeNftBoss()
            this.calBattleInfo()

            setTimeout(() => {
              this.showuptxtani = false
              this.showuptxtani2 = false
              this.showuptxtani3 = false
            }, 500)
          }
        }
      }
    },
    onPopClose() {
      this.isType = 0
      this.istip = false
      this.upgrading = false
    },
    calBattleInfo(){
      // 先初始化属性
      this.battleinfo.atk = 0
      this.battleinfo.hp = 0
      this.battleinfo.hit_rate = 0
      this.battleinfo.dodge_rate = 0
      this.battleinfo.cri_rate = 0
      this.battleinfo.cri_dmg_rate = 0
      this.battleinfo.atk_add_rate = 0
      this.battleinfo.hp_add_rate = 0
      this.battleinfo.atkAdd = 0
      this.battleinfo.hpAdd = 0
      // 先读人物基础属性
      this.battleinfo.atk += initAttr[0].BaseAtk
      this.battleinfo.hp += initAttr[0].BaseHP
      this.battleinfo.hit_rate += initAttr[0].HitRate
      this.battleinfo.dodge_rate += initAttr[0].DodgeRate
      this.battleinfo.cri_rate += initAttr[0].CriRate
      this.battleinfo.cri_dmg_rate += initAttr[0].CriDmgRate
      // 累加装备属性
      console.log(this.sbxData)
      for (let equip of this.sbxData) {
        this.battleinfo.atk += equip.attrs.atk;
        this.battleinfo.hp += equip.attrs.hp;
      }
      let addAtk = 0
      let addHp = 0
      // 计算NFT加成
      let nft1 = jskit.localStorageGet('equip1')
      let nft2 = jskit.localStorageGet('equip2')
      let nft3 = jskit.localStorageGet('equip3')
      let nftArr = []
      if(nft1){
        nftArr.push(nft1)
      }
      if(nft2){
        nftArr.push(nft2)
      }
      if(nft2){
        nftArr.push(nft3)
      }
      if(nftArr && nftArr.length > 0){
        for(let nft of nftArr){
          if(nft.add_atr_1_val > 0){
            if(nft.add_atr_1 == 'ATK'){
              addAtk += nft.add_atr_1_val
            }else if(nft.add_atr_1 == 'HP'){
              addHp += nft.add_atr_1_val
            }
          }
          if(nft.add_atr_2_val > 0){
            if(nft.add_atr_2 == 'ATK'){
              addAtk += nft.add_atr_2_val
            }else if(nft.add_atr_2 == 'HP'){
              addHp += nft.add_atr_2_val
            }
          }
          if(nft.main_atr_1 == 'ATK'){
            addAtk += nft.main_atr_1_val
          }else if(nft.main_atr_1 == 'HP'){
            addHp += nft.main_atr_1_val
          }else if(nft.main_atr_1 == 'HIT'){
            this.battleinfo.hit_rate += nft.main_atr_1_val
          }else if(nft.main_atr_1 == 'SPD'){
            this.battleinfo.dodge_rate += nft.main_atr_1_val
          }else if(nft.main_atr_1 == 'CRT'){
            this.battleinfo.cri_rate += nft.main_atr_1_val
          }
          if(nft.main_atr_2 == 'ATK'){
            addAtk += nft.main_atr_2_val
          }else if(nft.main_atr_2 == 'HP'){
            addHp += nft.main_atr_2_val
          }else if(nft.main_atr_2 == 'HIT'){
            this.battleinfo.hit_rate += nft.main_atr_2_val
          }else if(nft.main_atr_2 == 'SPD'){
            this.battleinfo.dodge_rate += nft.main_atr_2_val
          }else if(nft.main_atr_2 == 'CRT'){
            this.battleinfo.cri_rate += nft.main_atr_2_val
          }
        }
      }
      this.battleinfo.atk_add_rate = addAtk
      this.battleinfo.hp_add_rate = addHp
      this.battleinfo.atkAdd = Math.floor(this.battleinfo.atk * addAtk / 10000)
      this.battleinfo.hpAdd = Math.floor(this.battleinfo.hp * addHp / 10000)
      // 刷新session
      cache.setSession('battleInfo', this.battleinfo)
    },
    freshRedpoint(){
      this.redpoint.equipLv1 = false
      this.redpoint.equipLv2 = false
      this.redpoint.equipLv3 = false
      this.redpoint.equipLv4 = false
      this.redpoint.equipLv5 = false
      this.redpoint.equipStar1 = false
      this.redpoint.equipStar2 = false
      this.redpoint.equipStar3 = false
      this.redpoint.equipStar4 = false
      this.redpoint.equipStar5 = false
      for(let equip of this.sbxData){
        let lvNumNeed = this.baseEquipLevelData[equip.equip_type + "_" + equip.level].UpgradeConsumeItems[10]
        let curStar = this.baseEquipStarData[equip.equip_type + "_" + equip.star]
        let starNumNeed = 0
        if(equip.equip_type == 1){
          starNumNeed = curStar.UpgradeConsumeItems[11]
          if(this.materials.item10 >= lvNumNeed){
            this.redpoint.equipLv1 = true
          }
          if(this.materials.item11 >= starNumNeed){
            this.redpoint.equipStar1 = true
          }
        }else if(equip.equip_type == 2){
          starNumNeed = curStar.UpgradeConsumeItems[12]
          if(this.materials.item10 >= lvNumNeed){
            this.redpoint.equipLv2 = true
          }
          if(this.materials.item12 >= starNumNeed){
            this.redpoint.equipStar2 = true
          }
        }else if(equip.equip_type == 3){
          starNumNeed = curStar.UpgradeConsumeItems[13]
          if(this.materials.item10 >= lvNumNeed){
            this.redpoint.equipLv3 = true
          }
          if(this.materials.item13 >= starNumNeed){
            this.redpoint.equipStar3 = true
          }
        }else if(equip.equip_type == 4){
          starNumNeed = curStar.UpgradeConsumeItems[14]
          if(this.materials.item10 >= lvNumNeed){
            this.redpoint.equipLv4 = true
          }
          if(this.materials.item14 >= starNumNeed){
            this.redpoint.equipStar4 = true
          }
        }else if(equip.equip_type == 5){
          starNumNeed = curStar.UpgradeConsumeItems[15]
          if(this.materials.item10 >= lvNumNeed){
            this.redpoint.equipLv5 = true
          }
          if(this.materials.item15 >= starNumNeed){
            this.redpoint.equipStar5 = true
          }
        }
      }
      console.log("redpoint:",this.redpoint)
    },
    showUpAni(success){
      if(success){
        this.showupani = true
        let time = setTimeout(() => {
          this.showuptxtani = true
          clearTimeout(time)
        }, 500)
        let time2 = setTimeout(() => {
          this.showupani = false
          this.showuptxtani = false
          clearTimeout(time2)
        }, 1200)
      }
    },
    showUpgradingProgress(type){
      if(!this.upgrading){
        this.upgrading = true
      }
      this.upgradingVal = 0
      let perAdd = 4
      let tempVal = 0
      let handle = setInterval(() => {
        console.log(this.upgradingFinish)
        if(this.upgradingFinish){
          perAdd = 100
        }
        if(tempVal < 100){
          if(this.upgradingVal + perAdd > 100){
            this.upgradingVal = 100
            tempVal = 100
          }else{
            this.upgradingVal = this.upgradingVal + perAdd
            tempVal = tempVal + perAdd
          }
        }else{
          tempVal = 999
        }
        console.log("num", this.upgradingVal, tempVal)
        if(tempVal == 999){
          if(this.upgradingFinish){
            clearInterval(handle)
            if(this.upgradingSuccess){
              this.showUpAni(true)
            }else{
              if(type == 1){
                this.$refs.tip.openTip('Upgrade Failed.')
              }else if(type == 2){
                this.$refs.tip.openTip('Enchance Failed.')
              }
            }
            this.upgrading = false
            this.upgradingFinish = false
          }
        }
      },100)
    }
  }
}
</script>

<style lang='stylus' src="./index.css" />
