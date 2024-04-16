/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/05/29 9:46
 */

package logic

type group struct {
	User   user
	Email  email
	Movie  movie
	File   file
	Tag    tag
	Pm     pm
	Auto   auto
	Cinema cinema
	Plan   plan
	Seat   seat
	Ticket ticket
}

var Group = new(group)
