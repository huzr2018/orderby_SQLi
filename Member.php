<?php

class Controller_Member extends Controller {
 
    public function action_index() {
        $sort_by = $this->request->param('sort_by');
        $members = ORM::factory('Member');
        // Get all members order by $sort_by.
        $results = $members->order_by("username",$sort_by)->find_all();
        foreach($results as $member) {
            echo $member;
            echo '<br>';
        }
    }
}
?>
