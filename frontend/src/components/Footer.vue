<template>
  <footer class="footer-home ant-layout-footer" style="text-align: center;">
    <div class="div-footer">
      Sema - ©2020 Created by Team BlockChain Research / R&amp;D Sun - Asterisk
    </div>
  </footer>
</template>

<script>
export default {
  name: 'footer-page'
};
</script>

<style lang="scss">
.ant-layout-footer {
  padding: 0 50px 15px;
}

.footer-home {
  width: 100vw;
}
.ant-layout-footer {
  padding: 24px 50px;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  background: #f0f2f5;
}
.ant-layout-footer,
.ant-layout-header {
  -ms-flex: 0 0 auto;
  flex: 0 0 auto;
}

.footer-home {
  width: 100vw;
}

.div-footer {
  border-top: 1px solid gainsboro;
  text-align: center;
  padding: 1.5rem;
}

.ant-layout-footer {
  padding: 0px 50px 15px 50px;
  background: #f0f2f5;
}

@media all and (max-width: 768px) {
  .ant-layout-footer {
    padding: 0px 20px 15px 20px;
    background: #f0f2f5;
  }
}
</style>
