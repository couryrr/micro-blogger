# General Development Notes

## To start

I have been bouncing between windows and linux for development work. I really love linux and would really like to use it
when doing my web dev but I am an avid gamer as well. I do not want a device for gaming and a device for development.
I just want one.

So here I am developing on windows again.

Likewise, I have tried so hard to be cool and use nvim. Unfortunetly, I work for real companies not made up start-ups. So usually I am using
Java or C# (sometime python / go) and the setup is just too much. Also, why is a gui bad? I use the cli and know the underlying process of running
and debugging the code. So it is vscode/intelliJ for me.

Lastly, goodness do I have AWS tooling. So this is a project of first and I will be using Terraform!

## A bit more work

sam cli is actually working better than before. For the time being I am sticking with it.

I need to better understand the `template.yaml` mostly the output section

I would like to add an open api spec.

The api gateway should write to a queue for all but get.
get can be via pagaination with filter or id.

## Wow error messages in SAM suck

I took hours to figure out why permissions were not working. Turns out that there was 1 extra indent that caused it to not parse. You do not get a parsing error like "You cannot put that there" you just get a random permission error...

## Workflow

Not sure yet

## Terms

Stage - no idea
StageName - is used in the url... could be an env
Transform - A macro can be custom of aws provided.



Output - is what is shown at the end of you cli commnand and changeset output.
