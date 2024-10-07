<template>
  <div class="flex flex-col gap-2 font-sans">
    <!-- Image URL Input -->
    <div class="flex mt-1 gap-2">
      <input
        v-model="imageUrl"
        type="text"
        placeholder="Add image photo links"
        class="p-1 border w-4/5 border-solid border-gray-300 focus:outline-none rounded-sm"
      />
      <button
        @click.prevent="addImageLink"
        class="border p-2 bg-blue-500 text-white"
      >
        Add
      </button>
    </div>

    <!-- Display Image URLs -->
    <div class="mt-3 flex gap-5 flex-wrap">
      <div
        v-for="(image, index) in imageStore"
        :key="index"
        class="flex flex-col"
      >
        <img
          :src="image.image_url"
          class="rounded-md h-36 w-36 md:w-40 bg-cover bg-no-repeat"
          alt="Uploaded image"
        />
        <label class="font-serif text-center mt-1">Thumbnail:</label>
        <div class="flex justify-between">
          <label class="flex-item block">
            <input
              type="radio"
              :name="'thumbnail' + index"
              :value="true"
              @change="setImageThumbnail(index)"
              v-model="image.is_thumbnail"
            />
            Yes
          </label>
          <label class="flex-item block">
            <input
              type="radio"
              :name="'thumbnail' + index"
              :value="false"
              @change="setImageThumbnail(index)"
              v-model="image.is_thumbnail"
            />
            No
          </label>
          <i class="fa fa-trash mt-1" @click="removeImage(index)"></i>
        </div>
      </div>
    </div>

    <div class="mt-2">
      <input
        @change="uploadLocalImages"
        type="file"
        class="p-1 border w-4/5 border-solid border-gray-300 focus:outline-none rounded-sm"
        multiple
      />
    </div>

    <!-- Display Local Images -->
    <div v-if="localFiles.length > 0" class="mt-3 flex gap-5 flex-wrap">
      <div class="my-2 grid grid-cols-2 gap-5 md:grid-cols-3">
        <div
          v-for="(localFile, index) in localFiles"
          :key="localFile.staticPath"
          class="flex flex-col"
        >
          <img
            :src="localFile.staticPath"
            class="rounded-md h-36 w-36 md:w-40 bg-cover bg-no-repeat"
            alt="Local uploaded image"
          />
          <label class="font-serif text-center mt-1">Thumbnail:</label>
          <div class="flex justify-between">
            <label class="flex-item block">
              <input
                type="radio"
                :name="'localThumbnail' + index"
                :value="true"
                @change="setLocalThumbnail(index)"
                v-model="localFile.is_thumbnail"
              />
              Yes
            </label>
            <label class="flex-item block">
              <input
                type="radio"
                :name="'localThumbnail' + index"
                :value="false"
                @change="setLocalThumbnail(index)"
                v-model="localFile.is_thumbnail"
              />
              No
            </label>
            <i class="fa fa-trash mt-1" @click="removeLocalImage(index)"></i>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      imageUrl: "",
      imageStore: [],
      localFiles: [],
    };
  },
  methods: {
    addImageLink() {
      if (this.imageUrl.length > 0) {
        this.imageStore.push({ image_url: this.imageUrl, is_thumbnail: false });
        this.$emit("update-images", this.imageStore, this.localFiles); // Emit event
        this.imageUrl = "";
      }
    },
    removeImage(index) {
      this.imageStore.splice(index, 1);
      this.$emit("update-images", this.imageStore, this.localFiles); // Emit event
    },
    setImageThumbnail(index) {
      this.imageStore.forEach((image, i) => {
        image.is_thumbnail = i === index;
      });
      this.$emit("update-images", this.imageStore, this.localFiles); // Emit event
    },
    uploadLocalImages(e) {
      const localImageStore = [];
      for (let i = 0; i < e.target.files.length; i++) {
        const obj = {
          file: e.target.files[i],
          staticPath: URL.createObjectURL(e.target.files[i]),
          is_thumbnail: false,
        };
        localImageStore.push(obj);
      }
      this.localFiles = localImageStore;
      this.$emit("update-images", this.imageStore, this.localFiles); // Emit event
    },
    setLocalThumbnail(index) {
      this.localFiles.forEach((localFile, i) => {
        localFile.is_thumbnail = i === index;
      });
      this.$emit("update-images", this.imageStore, this.localFiles); // Emit event
    },
    removeLocalImage(index) {
      this.localFiles.splice(index, 1);
      this.$emit("update-images", this.imageStore, this.localFiles); // Emit event
    },
  },
};
</script>

<style scoped></style>
