

 <html>
    <head>
      <link rel="stylesheet" href="static/styles.css">

    </head>
    <body>
      <div class="home">
        <a href="http://localhost:8080/">Home</a>
      </div>
      <div class="header">
        <h1>Security Setup Page</h1>
        <p>Set and adjust Security Roles </p>
      </div>
      <div class="topnav">
        <a href="http://localhost:8080/">Home</a>
        <a href="http://localhost:8080/login">Login</a>
        <a href="#">Link</a>
        <a href="#">Link</a>
      </div>
      <div class="row">
        <div class="column side">
          <h2>left Side</h2>
          <p>Colum side row Lorem ipsum dolor sit amet, consectetur adipiscing elit..</p>
        </div>
        <div class="column middle">
          <h2>Main Content</h2>
          <p>Middle row: this is where you setup security for your organisation .</p>
          <p>Setup state city and school name and then the levels of permission you need </p>

          <form action="http://localhost:8080/postsecuritydata" method="post">
              <label> Organistaion name :  <input type="text" name="orgname" value="NSW Dept of Education"/></label>
              <label> Country :  <input type="text" name="country" value=  "AU" /></label>
              <label> State :  <input type="text" name="state" value='NSW'/></label>
              <label> City :  <input type="text" name="city" value="MossVale"/></label>
              <label> School Name :  <input type="text" name="sname" value="MossVale High"/></label>
              <label> Level 1 :  <input type="text" name="l1" value="Principal"/></label>
              <label> Level 2 :  <input type="text" name="l2" value="Head Teachers"/></label>
              <label> Level 3 :  <input type="text" name="l3" value="Teachers"/></label>
              <label> Level 4 :  <input type="text" name="l4" value="Students"/></label>
              <label> Level 5 :  <input type="text" name="l5" value="Students-Parent"/></label>
              <label> Level 6 :  <input type="text" name="l6" value=""/></label>
              <label> Level 7 :  <input type="text" name="l7" value=""/></label>
              <label> Level 8 :  <input type="text" name="l8" value=""/></label>
              <label> Level 9 :  <input type="text" name="l9" value=""/></label>

              <input type="submit" value="Save">
            </form>
        </div>
        <div class="column side">
          <h2>right side</h2>
          <p>last colum side row or sit amet, consectetur adipiscing elit..
          </p>
          <!-- commenting out text -->


        </div>
      </div>


    </body>
</html>
