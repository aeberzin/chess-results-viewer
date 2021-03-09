<template>
  <tiptap-vuetify
    v-model="content"
    :extensions="extensions"
    class="tt"
  />
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import {
  TiptapVuetify,
  Heading,
  Bold,
  Italic,
  Strike,
  Underline,
  Code,
  Paragraph,
  BulletList,
  OrderedList,
  ListItem,
  Link,
  Blockquote,
  HardBreak,
  HorizontalRule,
  Image,
  History
} from 'tiptap-vuetify';

@Component({
  components: { TiptapVuetify }
})
export default class Editor extends Vue {
  @Prop({ required: false }) value!: any;

  get content() {
    return this.value;
  }
  set content(value) {
    this.$emit('input', value);
  }

  private extensions: any = [
    History,
    Blockquote,
    Link,
    Underline,
    Strike,
    Italic,
    ListItem,
    BulletList,
    OrderedList,
    [
      Heading,
      {
        options: {
          levels: [1, 2, 3]
        }
      }
    ],
    Bold,
    Link,
    Code,
    HorizontalRule,
    Paragraph,
    HardBreak
  ];
}
</script>

<style lang="scss" scoped>
.tt {
  ::v-deep .ProseMirror {
    min-height: 200px;
  }
}
</style>