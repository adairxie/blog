{{define "footer"}}
    <!-- Footer -->
    <footer>
      <div class="container">
        <div class="row">
          <div class="col-lg-8 col-md-10 mx-auto">
            <ul class="list-inline text-center">
              <li class="list-inline-item">
                <a href="#">
                  <span class="fa-stack fa-lg">
                    <i class="fa fa-circle fa-stack-2x"></i>
                    <i class="fa fa-twitter fa-stack-1x fa-inverse"></i>
                  </span>
                </a>
              </li>
              <li class="list-inline-item">
                <a href="#">
                  <span class="fa-stack fa-lg">
                    <i class="fa fa-circle fa-stack-2x"></i>
                    <i class="fa fa-facebook fa-stack-1x fa-inverse"></i>
                  </span>
                </a>
              </li>
              <li class="list-inline-item">
                <a href="#">
                  <span class="fa-stack fa-lg">
                    <i class="fa fa-circle fa-stack-2x"></i>
                    <i class="fa fa-github fa-stack-1x fa-inverse"></i>
                  </span>
                </a>
              </li>
            </ul>
            <p class="copyright text-muted">Copyright &copy; Your Website 2017</p>
          </div>
        </div>
      </div>
    </footer>

    <!-- Bootstrap core JavaScript -->
    <script src="/static/vendor/jquery/jquery.min.js"></script>
    <script src="/static/vendor/bootstrap/js/bootstrap.bundle.min.js"></script>
    <script src="/static/locale/bootstrap-markdown.zh.js"></script>
    <script src="/static/locale/bootstrap-markdown.fr.js"></script>
    <script src="/static/js/bootstrap-markdown.js"></script>
    <script src="/static/js/markdown.js"></script>
    <script src="/static/js/to-markdown.js"></script>

    <!-- Custom scripts for this template -->
    <script src="/static/js/clean-blog.min.js"></script>
    <script type="text/javascript">
        //显示中文提示
        (function ($) {
            $.fn.markdown.messages.zh = {
                        'Bold': "粗体",
                        'Italic': "斜体",
                        'Heading': "标题",
                        'URL/Link': "链接",
                        'Image': "图片",
                        'List': "列表",
                        'Unordered List': "无序列表",
                        'Ordered List': "有序列表",
                        'Code': "代码",
                        'Quote': "引用",
                        'Preview': "预览",
                        'strong text': "粗体",
                        'emphasized text': "强调",
                        'heading text': "标题",
                        'enter link description here': "输入链接说明",
                        'Insert Hyperlink': "URL地址",
                        'enter image description here': "输入图片说明",
                        'Insert Image Hyperlink': "图片URL地址",
                        'enter image title here': "在这里输入图片标题",
                        'list text here': "这里是列表文本",
                        'code text here': "这里输入代码",
                        'quote here': "这里输入引用文本"
                      
            };
                
        }(jQuery));
         //初始化编辑器
         $("#editor").markdown({
                     autofocus: true,
                     language: 'zh',
                    
         })
     </script>

{{end}}
