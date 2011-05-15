<!--
	Copyright 2009 The Go Authors. All rights reserved.
	Use of this source code is governed by a BSD-style
	license that can be found in the LICENSE file.
-->

	<!-- PackageName is printed as title by the top-level template -->
		<p><code>import "."</code></p>
	
			<p>
			<h4>Package files</h4>
			<span style="font-size:90%">
				<a href="/">defs.go</a>
				<a href="/">git.go</a>
			</span>
			</p>
		<h2 id="Constants">Constants</h2>
			<p>
Constants
</p>

			<pre>const (
    GIT_OBJ_ANY       = -0x2
    GIT_OBJ_BAD       = -0x1
    GIT_OBJ__EXT1     = 0
    GIT_OBJ_BLOB      = 0x3
    GIT_OBJ_COMMIT    = 0x1
    GIT_OBJ_OFS_DELTA = 0x6
    GIT_OBJ_REF_DELTA = 0x7
    GIT_OBJ_TAG       = 0x4
    GIT_OBJ_TREE      = 0x2
    GIT_OBJ__EXT2     = 0x5
    GIT_SUCCESS       = 0
    GIT_ENOTOID       = -0x2
)</pre>
			
			<pre>const (
    NOTBARE = iota
    BARE
)</pre>
			<h2 id="CommitCreate">func <a href="/?s=1142:1244#L59">CommitCreate</a></h2>
			<p><code>func CommitCreate(repo *Repo, tree, parent *Oid, author, commiter *Signature, message string) os.Error</code></p>
			
			<h2 id="GetHeadString">func <a href="/?s=3035:3084#L153">GetHeadString</a></h2>
			<p><code>func GetHeadString(repo *Repo) (string, os.Error)</code></p>
			
			<h2 id="LastError">func <a href="/?s=5016:5041#L252">LastError</a></h2>
			<p><code>func LastError() os.Error</code></p>
			<p>
Helper functions
</p>

			<h2 id="Commit">type <a href="/?s=1092:1140#L55">Commit</a></h2>
			<p>
Commit
</p>

			<p><pre>type Commit struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre></p>
				<h3 id="Commit.Author">func (*Commit) <a href="/?s=1852:1884#L92">Author</a></h3>
				<p><code>func (c *Commit) Author() string</code></p>
				
				<h3 id="Commit.Email">func (*Commit) <a href="/?s=1956:1987#L96">Email</a></h3>
				<p><code>func (c *Commit) Email() string</code></p>
				
				<h3 id="Commit.Lookup">func (*Commit) <a href="/?s=1559:1614#L79">Lookup</a></h3>
				<p><code>func (c *Commit) Lookup(r *Repo, o *Oid) (err os.Error)</code></p>
				
				<h3 id="Commit.Msg">func (*Commit) <a href="/?s=1750:1779#L87">Msg</a></h3>
				<p><code>func (c *Commit) Msg() string</code></p>
				
			<h2 id="GitTime">type <a href="/?s=408:459#L15">GitTime</a></h2>
			
			<p><pre>type GitTime struct {
    Time   int64
    Offset int32
}</pre></p>
			<h2 id="Index">type <a href="/?s=3943:3988#L199">Index</a></h2>
			<p>
Index
</p>

			<p><pre>type Index struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre></p>
				<h3 id="Index.Add">func (*Index) <a href="/?s=4165:4212#L210">Add</a></h3>
				<p><code>func (v *Index) Add(file string) (err os.Error)</code></p>
				
				<h3 id="Index.Free">func (*Index) <a href="/?s=4646:4668#L233">Free</a></h3>
				<p><code>func (v *Index) Free()</code></p>
				
				<h3 id="Index.Open">func (*Index) <a href="/?s=3990:4037#L203">Open</a></h3>
				<p><code>func (v *Index) Open(repo *Repo) (err os.Error)</code></p>
				
				<h3 id="Index.Read">func (*Index) <a href="/?s=4360:4397#L219">Read</a></h3>
				<p><code>func (v *Index) Read() (err os.Error)</code></p>
				
				<h3 id="Index.Write">func (*Index) <a href="/?s=4502:4540#L226">Write</a></h3>
				<p><code>func (v *Index) Write() (err os.Error)</code></p>
				
			<h2 id="IndexEntry">type <a href="/?s=461:823#L20">IndexEntry</a></h2>
			
			<p><pre>type IndexEntry struct {
    Ctime          [12]byte <span class="comment">/* git_index_time */</span>
    Mtime          [12]byte <span class="comment">/* git_index_time */</span>
    Dev            uint32
    Ino            uint32
    Mode           uint32
    Uid            uint32
    Gid            uint32
    File_size      int64
    Oid            [20]byte <span class="comment">/* git_oid */</span>
    Flags          uint16
    Flags_extended uint16
    Path           *int8
}</pre></p>
			<h2 id="Oid">type <a href="/?s=2068:2107#L102">Oid</a></h2>
			<p>
Oid
</p>

			<p><pre>type Oid struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre></p>
				<h3 id="Oid.GetHead">func <a href="/?s=3245:3286#L163">GetHead</a></h3>
				<p><code>func GetHead(repo *Repo) (*Oid, os.Error)</code></p>
				
				<h3 id="Oid.NewOid">func <a href="/?s=2109:2127#L106">NewOid</a></h3>
				<p><code>func NewOid() *Oid</code></p>
				
				<h3 id="Oid.NewOidString">func <a href="/?s=2162:2206#L110">NewOidString</a></h3>
				<p><code>func NewOidString(s string) (*Oid, os.Error)</code></p>
				
				<h3 id="Oid.TreeFromIndex">func <a href="/?s=858:919#L45">TreeFromIndex</a></h3>
				<p><code>func TreeFromIndex(repo *Repo, index *Index) (*Oid, os.Error)</code></p>
				
				<h3 id="Oid.String">func (*Oid) <a href="/?s=2344:2373#L118">String</a></h3>
				<p><code>func (v *Oid) String() string</code></p>
				
			<h2 id="Reference">type <a href="/?s=3579:3636#L182">Reference</a></h2>
			<p>
Reference
</p>

			<p><pre>type Reference struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre></p>
				<h3 id="Reference.GetOid">func (*Reference) <a href="/?s=3845:3878#L194">GetOid</a></h3>
				<p><code>func (v *Reference) GetOid() *Oid</code></p>
				
				<h3 id="Reference.Lookup">func (*Reference) <a href="/?s=3638:3701#L186">Lookup</a></h3>
				<p><code>func (v *Reference) Lookup(r *Repo, name string) (err os.Error)</code></p>
				
			<h2 id="Repo">type <a href="/?s=264:312#L14">Repo</a></h2>
			<p>
Repo
</p>

			<p><pre>type Repo struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre></p>
				<h3 id="Repo.Free">func (*Repo) <a href="/?s=541:562#L28">Free</a></h3>
				<p><code>func (v *Repo) Free()</code></p>
				
				<h3 id="Repo.Init">func (*Repo) <a href="/?s=603:664#L32">Init</a></h3>
				<p><code>func (v *Repo) Init(path string, isbare uint8) (err os.Error)</code></p>
				
				<h3 id="Repo.Open">func (*Repo) <a href="/?s=314:361#L18">Open</a></h3>
				<p><code>func (v *Repo) Open(path string) (err os.Error)</code></p>
				
			<h2 id="RevWalk">type <a href="/?s=2489:2540#L126">RevWalk</a></h2>
			<p>
RevWalk
</p>

			<p><pre>type RevWalk struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre></p>
				<h3 id="RevWalk.NewRevWalk">func <a href="/?s=2542:2590#L130">NewRevWalk</a></h3>
				<p><code>func NewRevWalk(repo *Repo) (*RevWalk, os.Error)</code></p>
				
				<h3 id="RevWalk.Free">func (*RevWalk) <a href="/?s=3502:3526#L177">Free</a></h3>
				<p><code>func (v *RevWalk) Free()</code></p>
				
				<h3 id="RevWalk.Next">func (*RevWalk) <a href="/?s=2883:2928#L146">Next</a></h3>
				<p><code>func (v *RevWalk) Next(o *Oid) (err os.Error)</code></p>
				
				<h3 id="RevWalk.Push">func (*RevWalk) <a href="/?s=2801:2831#L142">Push</a></h3>
				<p><code>func (v *RevWalk) Push(o *Oid)</code></p>
				
				<h3 id="RevWalk.Reset">func (*RevWalk) <a href="/?s=2734:2759#L138">Reset</a></h3>
				<p><code>func (v *RevWalk) Reset()</code></p>
				
				<h3 id="RevWalk.Sorting">func (*RevWalk) <a href="/?s=3462:3496#L174">Sorting</a></h3>
				<p><code>func (v *RevWalk) Sorting(sm uint)</code></p>
				<p>
TODO: implement this
</p>

			<h2 id="Signature">type <a href="/?s=4718:4775#L238">Signature</a></h2>
			<p>
Signature
</p>

			<p><pre>type Signature struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre></p>
				<h3 id="Signature.NewSignature">func <a href="/?s=4777:4825#L242">NewSignature</a></h3>
				<p><code>func NewSignature(name, email string) *Signature</code></p>
				
			<h2 id="Test">type <a href="/?s=227:254#L11">Test</a></h2>
			
			<p><pre>type Test *C.git_repository</pre></p>
			<h2 id="Tree">type <a href="/?s=814:856#L41">Tree</a></h2>
			<p>
Tree
</p>

			<p><pre>type Tree struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre></p>
