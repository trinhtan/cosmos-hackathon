<template>
  <div>
    <div class="content-new">
      <div class="box-shadow">
        <el-steps :space="200" :active="current" simple>
          <el-step icon="el-icon-edit"></el-step>
          <el-step icon="el-icon-picture"></el-step>
          <el-step icon="el-icon-info"></el-step>
        </el-steps>
        <el-divider content-position="center"><h3>Fill in the information</h3></el-divider>

        <div :class="current === 0 ? 'steps-content' : 'steps-content display-none'" id="step-0">
          <div class="ant-modal-body">
            <el-form
              :label-position="labelPosition"
              :model="form1"
              :rules="rules_1"
              ref="ruleForm_1"
            >
              <el-form-item label="Title" prop="title">
                <el-input v-model="form1.title" placeholder="Please input title..."></el-input>
              </el-form-item>
              <el-form-item label="Category" prop="category">
                <el-input
                  v-model="form1.category"
                  placeholder="Please input catagory..."
                ></el-input>
              </el-form-item>
            </el-form>
          </div>
        </div>
        <div :class="current === 1 ? 'steps-content' : 'steps-content display-none'" id="step-1">
          <div class="ant-modal-body">
            <vue-upload-multiple-image
              @upload-success="selectImageSuccess"
              @before-remove="beforeRemove"
              @edit-image="editImage"
              :class="'box-upload-image'"
              browseText=""
              dragText="select(or drag) the image here"
              primaryText=""
              :showPrimary="false"
              dropText="Drag your files here..."
            >
            </vue-upload-multiple-image>
            <div class="error-form-2" v-if="!rules_2">
              Please select image
            </div>
          </div>
        </div>
        <div :class="current === 2 ? 'steps-content' : 'steps-content display-none'" id="step-2">
          <div class="ant-modal-body">
            <el-form
              :label-position="labelPosition"
              :model="form3"
              :rules="rules_3"
              ref="ruleForm_3"
            >
              <el-form-item label="Description" prop="description">
                <el-input
                  type="textarea"
                  :rows="5"
                  placeholder="Please input description..."
                  v-model="form3.description"
                >
                </el-input>
              </el-form-item>
            </el-form>
          </div>
        </div>

        <el-button
          :disabled="current !== 0 ? false : true"
          style="margin-top: 12px;"
          size="small"
          round
          @click="prev"
          ><i class="el-icon-arrow-left"></i> Prev</el-button
        >
        <el-button
          v-if="current !== 2"
          style="margin-top: 12px;"
          size="small"
          round
          @click="next"
          type="primary"
          >Next <i class="el-icon-arrow-right"></i
        ></el-button>
        <el-button
          v-if="current === 2"
          style="margin-top: 12px;"
          size="small"
          round
          @click="publishData"
          type="primary"
          v-loading.fullscreen.lock="loadingUpload"
          >Publish <i class="el-icon-upload el-icon-right"></i
        ></el-button>
      </div>
    </div>
  </div>
</template>

<script>
import VueUploadMultipleImage from '../../node_modules/vue-upload-multiple-image';
import { ipfsNodeUri } from '@/utils/config.js';
import getIpfs from '@/utils/getIpfs';
import { streamFiles } from '@/utils/checkFileUpload.js';
import { mapState, mapActions } from 'vuex';
export default {
  name: 'publish',
  data() {
    return {
      current: 0,
      labelPosition: 'top',
      formAsset: {
        title: '',
        category: '',
        description: ''
      },
      fileList: [],
      form1: {
        title: '',
        category: ''
      },
      rules_1: {
        title: [{ required: true, message: 'Please input title', trigger: 'blur' }],
        category: [{ required: true, message: 'Please input catagory', trigger: 'blur' }]
      },
      form2_images: [],
      rules_2: true,
      form3: {
        description: ''
      },
      rules_3: {
        description: [{ required: true, message: 'Please input description', trigger: 'blur' }]
      },
      loadingUpload: false
    };
  },
  computed: {
    ...mapState('cosmos', ['products'])
  },
  components: {
    VueUploadMultipleImage
  },

  methods: {
    ...mapActions('cosmos', ['createProduct', 'signTxt', 'getAllProducts', 'setCosmosAccount']),
    next() {
      if (this.current === 0) {
        this.$refs['ruleForm_1'].validate((valid) => {
          if (valid) {
            this.formAsset.title = this.form1.title;
            this.formAsset.category = this.form1.category;
            if (this.current++ > 2) this.current = 2;
          } else {
            return false;
          }
        });
      } else if (this.current === 1) {
        if (this.form2_images.length > 0) {
          this.rules_2 = true;
          if (this.current++ > 2) this.current = 2;
        } else {
          this.rules_2 = false;
          return false;
        }
      }
    },
    prev() {
      if (this.current-- < 0) this.current = 0;
    },
    selectImageSuccess(formData, index, fileList) {
      this.form2_images.push(formData.get('file'));
      this.rules_2 = true;
      this.fileList = fileList;
    },
    beforeRemove(index, done, fileList) {
      done();
      this.form2_images.splice(index, 1);
      this.fileList = fileList;
    },
    editImage(formData, index, fileList) {
      this.form2_images[index] = formData.get('file');
      this.fileList = fileList;
    },
    publishData() {
      this.$refs['ruleForm_3'].validate(async (valid) => {
        if (valid) {
          this.formAsset.description = this.form3.description;
          this.loadingUpload = true;
          let images = [];
          for (let i = 0; i < this.form2_images.length; i++) {
            let file = this.form2_images[i];
            var url = await this.uploadImageIpfs(file);
            images.push(url);
            console.log(url);
          }
          var reponseCreate = await this.createProduct({ asset: this.formAsset, images });
          await this.setCosmosAccount();
          this.loadingUpload = false;
          await this.open(reponseCreate);
        } else {
          return false;
        }
      });
    },
    async addToIpfs(data) {
      try {
        const { hostname, port, protocol } = new URL(ipfsNodeUri);
        const ipfsConfig = {
          protocol: protocol.replace(':', ''),
          host: hostname,
          port: port || '443'
        };
        const { ipfs } = await getIpfs(ipfsConfig);
        const cid = await streamFiles(ipfs, data);
        return cid;
      } catch (error) {
        console.log(error);
      }
    },
    async uploadImageIpfs(e) {
      let dataFile = e;
      let data_cid = await this.addToIpfs({ path: dataFile.name, content: dataFile });
      let linkFile = `https://gateway.ipfs.io/ipfs/${data_cid}/${dataFile.name}`;
      if (linkFile) {
        return linkFile;
      } else {
        console.log('error upload ipfs');
      }
    },
    open(reponseCreate) {
      const h = this.$createElement;
      this.$msgbox({
        title: 'Confirm',
        message: h('p', null, [h('span', null, 'Do you want to broadcast the transaction?')]),
        showCancelButton: true,
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        beforeClose: (action, instance, done) => {
          if (action === 'confirm') {
            instance.confirmButtonLoading = true;
            instance.confirmButtonText = 'Loading...';
            done();
            instance.confirmButtonLoading = false;
          } else {
            done();
          }
        }
      }).then(async () => {
        this.loadingUpload = true;
        var reponseSign = await this.signTxt(reponseCreate);
        console.log(reponseSign);
        await setTimeout(async () => {
          await this.getAllProducts();
          await this.$message({
            type: 'success',
            message: 'Publish success'
          });
          this.loadingUpload = false;
          await this.$router.push({ name: 'MyAssets' });
        }, 5000);
      });
    }
  }
};
</script>

<style lang="scss">
.content-new {
  padding-top: 170px;
  width: 40%;
  margin: 0 auto;
  min-height: 900px;
}
.box-shadow {
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 10px 16px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}

.el-step__title.is-process {
  font-weight: 435;
  color: #3a8ee6;
}
.el-step__head.is-process {
  color: #2c75ff;
  border-color: #2c75ff;
}

.steps-content {
  margin-top: 20px;
  border: 1px dashed #e9e9e9;
  border-radius: 6px;
  background-color: #fafafa;
  min-height: 200px;
  text-align: center;
}

.display-none {
  display: none;
}

.ant-modal-body {
  padding: 24px;
  font-size: 14px;
  line-height: 1.5;
  word-wrap: break-word;
}

.steps-content {
  label.el-form-item__label {
    display: block;
  }
  div.el-input-number {
    width: 100%;
    input {
      text-align: left;
    }
  }
}

.box-upload-image {
  .image-container {
    margin: 0 auto;
  }
  .image-list-container {
    margin: 0 auto;
    margin-top: 10px;
    .image-list-item {
      height: 50px;
      width: 58px;
      .show-img {
        max-width: 50px;
        max-height: 50px;
      }
    }
  }
  div.show-image.centered {
    padding: 5px;
    width: 100%;
    img {
      width: 100%;
      max-width: 100%;
      max-height: 100%;
    }
  }
}

.error-form-2 {
  margin-top: 10px;
  color: #f56c6c;
  font-size: 12px;
  line-height: 1;
  padding-top: 4px;
  top: 100%;
  left: 0;
}
@media (min-width: 1024px) {
  .box-upload-image {
    .image-container {
      width: 335px;
      height: 240px;
      .preview-image {
        .show-image {
          top: 70%;
        }
        .image-overlay {
          height: 185px;
        }
        .image-overlay-details {
          top: 70%;
        }
      }
    }
    .image-list-container {
      max-width: 330px;
      .image-list-item {
        width: 62px;
      }
    }
  }
}

@media (max-width: 768px) {
  .content-new {
    padding-top: 170px;
    width: 85%;
    margin: 0 auto;
  }
}
</style>
