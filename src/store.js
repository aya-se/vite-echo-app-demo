import { computed, reactive, ref } from "vue";

export const store = reactive({
  count: 0,
  globalList: ["りんご", "みかん", "バナナ"],
});

// ゲームの状態を管理する
export const status = ref("title");

// タイマーを管理する
export const timer = reactive({
  countDownTimer: 3,
  gameTimer: 30,
});

// 得点を管理する
export const score = reactive({
  ika: 0,
  shika: 0,
  meka: 0,
  total: computed(() => {
    return score.ika + score.shika + score.meka;
  }),
  highest: 0,
});

// メッセージリストを管理する
export const messages = reactive({
  allList: [],
  correctList: computed(() => {
    return messages.allList.filter((message) => {
      return message.ika || message.shika || message.meka;
    });
  }),
  incorrectList: computed(() => {
    return messages.allList.filter((message) => {
      return !message.ika && !message.shika && !message.meka;
    });
  }),

  // ゲーム画面用
  leftList: [],
  middleList: [],
  rightList: [],

  // リザルト画面用
  resultIkaList: [],
  resultShikaList: [],
  resultMekaList: [],
});

// ランキングを管理する
export const ranking = reactive({
  list: [],
});
